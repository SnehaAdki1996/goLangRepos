package models

type User struct {
	userId int
	f_name string
	l_name string
}

var (
	users  []*User //Slice of pointer
	nextID int32   = 1
)
