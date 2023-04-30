package main

import (
	"fmt"
)

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	var arr1 = [...]int{1, 2, 3}
	fmt.Println(myArray)
	fmt.Println(arr1)
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}
