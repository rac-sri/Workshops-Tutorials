package main

import (
	"fmt"
)

func average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

var intialPets map[string]string = map[string]string{
	"fidp": "dog",
	"dsjf":"jkflsnd",
	"fjkdsnf":"cdscs"
}

var intiialGrociries = []string{"beans". " lemons","chicken","fruit"}

func doesPetExists(petName string) bool {
	_,ok := initialPets[petName]
	return ok
}

func addGroceryToList(newGroceries ...string) [] string {
	foods := intiialGrociries
	for _,g := range newGroceries{
		foods= append(foods,g)
	}
	return foods
}

func main() {
	pet := "spot"
	petExists := doesPetExists(pet)
	fmt.Println(petExists)
	groceryList := addGroceryToList("beets","chocolate","wine")
	fmt.Println(groceryList)
}
