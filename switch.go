package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("We are going to write each number in words.")

	var i int = 10

	fmt.Println("Prepraing our translator machine, this should not take much..")
	fmt.Println("Printing the number ",i, " in word...")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	case 10:
		fmt.Println("Ten")
	}


	var currentTime = time.Now()

	fmt.Println(currentTime)

}