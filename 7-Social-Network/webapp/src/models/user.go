package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Nick         string        `json:"nick"`
	CreationDate time.Time     `json:"creationDate"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

func GetCompleteUser(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationsChannel := make(chan []Publication)

	go GetUserData(userChannel, userID, r)
	go GetFollowers(followersChannel, userID, r)
	go GetFollowing(followingChannel, userID, r)
	go GetPublications(publicationsChannel, userID, r)

	var (
		user         User
		followers    []User
		following    []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case loadedUser := <-userChannel:
			if loadedUser.ID == 0 {
				return User{}, errors.New("error when getting user")
			}

			user = loadedUser

		case loadedFollowers := <-followersChannel:
			if loadedFollowers == nil {
				return User{}, errors.New("error when getting followers")
			}

			followers = loadedFollowers

		case loadedFollowing := <-followingChannel:
			if loadedFollowing == nil {
				return User{}, errors.New("error when getting following")
			}

			following = loadedFollowing

		case loadedPublications := <-publicationsChannel:
			if loadedPublications == nil {
				return User{}, errors.New("error when getting publications")
			}

			publications = loadedPublications
		}
	}

	user.Followers = followers
	user.Following = following
	user.Publications = publications

	return user, nil
}

func GetUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func GetFollowers(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)

	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User

	err = json.NewDecoder(response.Body).Decode(&followers)
	if err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

func GetFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)

	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User

	err = json.NewDecoder(response.Body).Decode(&following)
	if err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

func GetPublications(channel chan<- []Publication, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.APIURL, userID)

	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var publications []Publication

	err = json.NewDecoder(response.Body).Decode(&publications)
	if err != nil {
		channel <- nil
		return
	}

	if publications == nil {
		channel <- make([]Publication, 0)
		return
	}

	channel <- publications
}
