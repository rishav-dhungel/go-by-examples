package main

import "fmt"

func main() {
	var i int = 1
	for i <= 10 {
		fmt.Println("Go is coolest programming language,", i)
		i++
	}

	var j int = 100
	for ; j>0; j-- {
		fmt.Println("We are reducing number values......from 100 a step at a time", j)
	}

	for n := range 10{
		fmt.Println(n)
	}

}