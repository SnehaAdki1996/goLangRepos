package main

import (
	"net/http"

	//"github.com:SnehaAdki1996/goLangRepo/models"
	"fmt"

	"main.go/controller"
	//"github.com/SnehaAdki1996/goLangRepos/controller"
)

func main() {

	//Declare and Asssign
	var i int
	i = 10
	fmt.Println(i)

	//Declare a time of initalization
	var f float32 = 2.3
	fmt.Println(f)

	//Dynamic typecasting
	firstName := "Sneha Vijay Konkati"
	fmt.Println(firstName)
	fmt.Println("Hello World")

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	//Working with Pointers
	var fname *string = new(string)
	fmt.Println(fname) //will print the Adress 0xc00004c1d0
	*fname = "Sneha"
	*fname = "Vijay" //Pointers can be reinitialized
	fmt.Println(*fname, fname)

	ename := "Adki"
	ptr := &ename
	fmt.Println(ptr, *ptr)

	ename = "Konkati"
	fmt.Println(ptr, *ptr)

	const c1 = 1
	fmt.Println(c1 + 4)
	fmt.Println(c1 + 4.5)

	const c3 int = 4
	fmt.Println(c3 + 4)
	fmt.Println(float32(c3) + 1.4)

	controller.RegisterController()
	http.ListenAndServe(":3000", nil)
}

// C:\Go_Code>go run main.go
// 10
// 2.3
// Sneha Vijay Konkati
// Hello World
// true
// (3+4i)
// 3 4
// 0xc0000401e0
// Vijay 0xc0000401e0
// 0xc00004c1f0 Adki
// 0xc00004c1f0 Konkati
// 5
// 5.5
// 8
// 5.4
