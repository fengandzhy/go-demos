package main

import "fmt"

type ABC int
func main() {
	var x = 20
	const y = 30
	fmt.Println(x)
	fmt.Println(y)
	var a ABC = 23
	fmt.Println(a)
}
