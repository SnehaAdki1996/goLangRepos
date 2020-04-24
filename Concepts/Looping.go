package main

func main() {
	var i int

	//Loop till condition
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			//break //will exit completely
			continue
		}

		println("Contining....")
	}

	//POST Clause with For Loop
	println("POST Clause with For Loop")
	for j := 0; j < 5; j++ {
		println(j)
	}

	//INFINITE LOOP
	println("Infinite")
	//ugle way
	var k int
	for {
		if k == 5 {
			break
		}
		println(k)
		k++
	}

	//loop over collection
	//1st way by len()
	slice := []int{10, 20, 30}
	for i := 0; i < len(slice); i++ {
		println("Collection o/p by for loop :", slice[i])
	}

	//2nd way
	for i, v := range slice {
		println("By Range", i, v)
	}

	//collection using Map
	mostCommonHHTPPOrts := map[string]int{
		"http":  80,
		"https": 443,
	}

	///looping the map to get key and value
	for k, v := range mostCommonHHTPPOrts {
		println("Map Key ", k, " and Value ", v)
	}

	//to get only keys
	for k := range mostCommonHHTPPOrts {
		println("Only Keys ", k)
	}

	//Only values
	for _, k := range mostCommonHHTPPOrts {
		println("only values ", k)
	}
}

//OUTPUT
// C:\Go_Code\Concepts>go run Looping.go
// 0
// Contining....
// 1
// Contining....
// Contining....
// 1
// Contining....
// 2
// 3
// Contining....
// 4
// // Contining...
// POST Clause with For Loop
// 0
// 1
// 2
// 3
// 4
// Infinite
// 0
// 1
// 2
// 3
// 4
// Collection o/p by for loop : 10
// Collection o/p by for loop : 20
// Collection o/p by for loop : 30
// By Range 0 10
// By Range 1 20
// By Range 2 30
// Map Key  http  and Value  80
// Map Key  https  and Value  443
// Only Keys  http
// Only Keys  https
// only values  80
// only values  443
