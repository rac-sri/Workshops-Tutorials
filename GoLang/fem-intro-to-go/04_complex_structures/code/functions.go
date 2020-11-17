package main

import "fmt"

func printAge(age int) (ageofsally int) {
	for _, value := range ages {
		fmt.Println(value)
	}
	ageofsally = 21
	fmt.Println(age)
	return 
}

func main() {
	age1, age2 := printAge()
	fmt.Println(printAge(18))
}
