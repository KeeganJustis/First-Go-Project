package main

import (
	"fmt"
	"strconv"
)

func listofnumbers(x int) []int {
	var theslice []int
	for i := 0; i <= x; i++ {
		theslice = append(theslice, i)

	}
	return theslice
}

func evenOrOdd(y []int) []string {
	var finalslice []string
	for i := range y {
		if i%2 == 0 {
			finalslice = append(finalslice, strconv.Itoa(y[i])+" is even")
		} else {
			finalslice = append(finalslice, strconv.Itoa(y[i])+" is odd")
		}

	}
	return finalslice
}

func main() {

	results := evenOrOdd(listofnumbers(10))
	//fmt.Println(evenOrOdd(listofnumbers(10)))
	for index := range results {
		fmt.Println(results[index])
	}

}
