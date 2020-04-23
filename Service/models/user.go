package models

type User struct {
	userId int
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

func AddUSer(u User) (User, error) {
	u.userId = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
