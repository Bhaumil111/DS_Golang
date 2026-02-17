package main

import "fmt"

// const A = 1e1000
//const S = "hello" + "world " // this is constant folding as this concatenation happens in compile time means here only

// const A = len(S)
var x = 10

// y :=10
func main() {
	//
	//start := time.Now()
	//var s string
	//fmt.Println()
	//for i := 0; i < 50000; i++ {
	//
	//	s += "a"
	//}
	//
	//t1 := time.Since(start)
	//
	//fmt.Printf("For string %v\n", t1)
	//s2 := time.Now()
	//var S strings.Builder
	//fmt.Println(S)
	//for i := 0; i < 50000; i++ {
	//	S.WriteString("a")
	//}
	//
	//t2 := time.Since(s2)
	//
	//fmt.Printf("For stringbuilder %v\n", t2)

	s := "你好"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

}
