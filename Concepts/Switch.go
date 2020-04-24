package main

//fallthrough keyword is used to continue to next case
//no break keyword is used
func main() {
	Method := "GET"
	switch Method {
	case "GET":
		println("GET Request")
	case "POST":
		println("POST Request")
	case "DELETE":
		println("DELETE Request")
	case "PUT":
		println("PUT Request")
	default:
		println("Unhandeled")
	}
}

// C:\Go_Code\Concepts>go run Switch.go
// GET Request
