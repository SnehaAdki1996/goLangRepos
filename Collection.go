package main

import (
	"fmt"
)

func main() {
	//2 ways to declare an array

	//1st
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	//Declare with initialize
	arr2 := [3]int{2, 3, 4}
	fmt.Println(arr2)

	//Slice
	//Dynamic Array
	slice := []int{1, 2, 3}
	fmt.Println("Slice Value", slice)

	//we can append one value
	slice = append(slice, 4)
	fmt.Println("Slice Value After Single Append", slice)

	//we can also append multiple values
	slice = append(slice, 4, 27, 90)
	fmt.Println("Slice Value After Multiple Append", slice)

	//slice of slice
	s2 := slice[:]
	fmt.Println("Slice s2: ", s2)
	s3 := slice[1:3]
	fmt.Println("Slice s3 start index is 1 end index is 3: ", s3)
	s4 := slice[1:]
	fmt.Println("Slice s3 start index is 1 ", s4)
	s5 := slice[:4]
	fmt.Println("slice s5 with end Index 4", s5)

	//MAP
	m := map[string]int{"foo": 43}
	fmt.Println(m)
	fmt.Println(m["foo"])

	//we can chnage the Key Value of map anytime
	m["foo"] = 48
	fmt.Println(m)

	//the map provides builtin Delete =>{1st parameter is map name 2ndParameter is key}
	delete(m, "foo")
	fmt.Println(m)

	//struct
	//Declare struct
	type user struct {
		ID    int
		Fname string
		Lname string
	}

	//Long way to use struct and assign values
	var u user
	fmt.Println(u)
	u.ID = 1
	u.Fname = "Sneha"
	u.Lname = "Konkati"
	fmt.Println(u)

	//shortest way to use and Multiline assigneshould alwasy end with comma
	u2 := user{
		ID:    2,
		Fname: "Vijay",
		Lname: "Konkati",
	}
	fmt.Println(u2)
}

// OUTPUT:-
// C:\Go_Code>go run Collection.go
// [1 2 3]
// [2 3 4]
// Slice Value [1 2 3]
// Slice Value After Single Append [1 2 3 4]
// Slice Value After Multiple Append [1 2 3 4 4 27 90]
// Slice s2:  [1 2 3 4 4 27 90]
// Slice s3 start index is 1 end index is 3:  [2 3]
// Slice s3 start index is 1  [2 3 4 4 27 90]
// slice s5 with end Index 4 [1 2 3 4]
// map[foo:43]
// 43
// map[foo:48]
// map[]
// {0  }
// {1 Sneha Konkati}
// {2 Vijay Konkati}
