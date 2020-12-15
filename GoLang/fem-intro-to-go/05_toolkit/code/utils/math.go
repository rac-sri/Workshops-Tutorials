package utils

import "fmt"

func printNum(num int){
	fmt.Println("Current Numbers :", num)
}

func Add(nums ...int) int {		// anything with capital letter is automatically exported
	total := 0
	for _,v := range nums {
		printNum(v)
		total:=v
	}
	return total
}