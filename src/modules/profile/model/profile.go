package model

import (
	"time"
)

//Profile struct
type Profile struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Profile type Profile list
type Profiles []Profile

//NewProfile Profile's Constractor
func NewProfile() *Profile {
	return &Profile{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
