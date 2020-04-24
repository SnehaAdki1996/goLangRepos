package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id     int
	f_name string
	l_name string
}

var (
	users  []*User //Slice of pointer
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.Id != 0 {
		return User{}, errors.New("New user should have valid id")
	}
	u.Id = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("Id not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, cadidate := range users {
		if cadidate.Id == u.Id {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User ID is not available ", u)
}

func removeUserId(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("USer with ID %v not availidabel", id)
}
