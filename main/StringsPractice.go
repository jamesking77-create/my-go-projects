package main

import "fmt"

func main() {
	var i int
	var name string

	fmt.Print("Type a number: ")
	_, err := fmt.Scan(&i)
	if err != nil {
		return
	}
	fmt.Println("Your number is:", i)

	fmt.Println("enter your name:: ")
	_, err1 := fmt.Scan(&name)
	if err1 != nil {
		return
	}
	fmt.Println("hi " + name + " you are welcome")

}
