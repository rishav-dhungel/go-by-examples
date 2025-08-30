package main

import "fmt"

func plus(a int, b int, c string) int{
	return a + b
}

func multiple(a, b, c int) int{
	return a + b + c
}

func main() {
	fmt.Println(plus(1,2, "hello"))
	fmt.Println(multiple(1,2,3))

}

