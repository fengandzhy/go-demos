package main

import "fmt"

func main() {
	fmt.Println(fib(12))
	names:=[]string{"Bill","Mike","John"}
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
}

func fib(n int) int{
	x,y:=0,1
	for i:=0;i<n;i++{
		x,y = y,x+y
	}
	return x
}
