package main

import "fmt"

func main() {
	mySentence := "Hello this is a sentence." // just to drop var keyword
	var name string 
	
	for _, value := range mySentence { // _ allows to run code without using the variable
		fmt.Println(string(value))
	}
}
