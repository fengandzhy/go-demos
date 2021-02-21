package main

import (
	"fmt"
	"strconv"
)

type Year uint16
type Month uint8
type Day uint8
type Salary float32
type Flag bool

func printDate(year Year,month Month,day Day){
	fmt.Println(strconv.Itoa(int(year))+"年"+strconv.Itoa(int(month))+"月"+strconv.Itoa(int(day))+"日")
}

func main() {
	var year Year = 2018
	var month Month = 1
	var day Day = 12
	printDate(year,month,day)
	var salary Salary = 300.12
	var flag Flag = true
	if flag {
		fmt.Println(salary)
	}
}
