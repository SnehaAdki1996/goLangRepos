package main

import (
	"fmt"
)

//multiple const declaration

const (
	first   = "One"
	secound = "Second"
)

//Declare const outside
const (
	PI = 3.14
)

//iota (by default iota has 0)
//iota will increment by default
//iota is reset in every const block
//on iota we can apply Arithemtic operations
const (
	frollNumber = iota
	funassigned
	fRoll = iota + 2
)

const (
	srollNumber = iota
)

func main() {
	//wither inside main its accessible in main ya outside main at import level
	fmt.Println("Const at global Declaraing and Assigning", PI)
	fmt.Println("Multi Value Declare in const", first, secound)
	fmt.Println("IOTA:", frollNumber)
	fmt.Println("IOTA on increment", fRoll)
	fmt.Println("IOTA on unsassigned", funassigned)
	fmt.Println("REser on every Const block", srollNumber)
}

// //OUTPUT:
// Const at global Declaraing and Assigning 3.14
// Multi Value Declare in const One Second
// IOTA: 0
// IOTA on increment 4
// IOTA on unsassigned 1
// REser on every Const block 0
