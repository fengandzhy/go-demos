package main

import "fmt"

func main() {
	p:= new(int)
	fmt.Println(*p)
	*p = 123
	fmt.Println(*p)
}

func newInt() *int{
	return new(int)
}

func newInt1() *int{
	var value int
	return &value
}
