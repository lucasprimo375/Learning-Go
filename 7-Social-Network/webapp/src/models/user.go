package models

import (
	"net/http"
	"time"
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
}

func GetUserData(channel <-chan User, userID uint64, r *http.Request) {

}

func GetFollowers(channel <-chan []User, userID uint64, r *http.Request) {

}

func GetFollowing(channel <-chan []User, userID uint64, r *http.Request) {

}

func GetPublications(channel <-chan []Publication, userID uint64, r *http.Request) {

}
