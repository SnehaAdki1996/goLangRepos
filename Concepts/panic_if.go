package main

type User struct {
	Id    int
	fname string
	lname string
}

//else block in if else ya else if is not compulsory
func main() {
	u1 := User{
		Id:    1,
		fname: "Sneha",
		lname: "Konkati",
	}
	u2 := User{
		Id:    2,
		fname: "Viijay",
		lname: "Konkati",
	}
	if u1.Id == u2.Id {
		println("Id's are Same")
	} else if u1.lname == u2.lname {
		println("Last Name are same")
	} else {
		println("Different")
	}
}

// C:\Go_Code\Concepts>go run  panic_if.go
// Last Name are same
