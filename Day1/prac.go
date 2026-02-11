package main

import "fmt"

func main() {
	//fmt.Println("Learning about fixed size arrays")
	//var s [8][8]int
	//s[2][3] = 1
	////println(s)
	//fmt.Println(s)

	fmt.Println("learning about twod slices")

	var rows = 10
	var cols = 5

	slice := make([][]int, rows)

	for i := 0; i < rows; i++ {
		slice[i] = make([]int, cols)

	}

	fmt.Println(slice)
	println(slice)

	// Applications of crytographic has
	var hash [32]byte
	hash[0] = 1
	hash[1] = 'a'
	hash[2] = 'A'
	//hash[3] = 256

	fmt.Println(hash)
}
