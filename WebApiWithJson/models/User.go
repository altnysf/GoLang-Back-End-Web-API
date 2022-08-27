package models

type User struct {
	ID int
	Username string
	Firstname string
	Lastname string
	Profile string
	Interests []Interest
}