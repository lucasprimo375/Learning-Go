package models

import "time"

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
