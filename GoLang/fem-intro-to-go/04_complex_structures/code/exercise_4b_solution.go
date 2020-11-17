package main

import "fmt"

func average ( numbers ...float64) float64 {
	 // unknown number of flaot64 numbers

	 total:=0.0

	 for _, number := range numbers {
		 total +=number
	 }

	 return total/float64(len(numbers))
}

func main(){
	fmt.Println(average(10,5,7))
}