package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func printName() func() string {
	newString := "-"
	return func() string{
		newString +=  "*"
		return newString
	}
}



func main() {
	nextInt := intSeq()
	nextNextInt := intSeq()
	nextPrint := printName()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("==========")
	fmt.Println(nextNextInt())
	fmt.Println(nextNextInt())
	fmt.Println(nextNextInt())
	fmt.Println("==========")
	fmt.Println(nextPrint())
	fmt.Println(nextPrint())
	fmt.Println(nextPrint())
	fmt.Println(nextPrint())
	fmt.Println(nextPrint())
	fmt.Println(nextPrint())
	fmt.Println("==========")


}