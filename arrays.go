package main

import "fmt"

func main() {

	var a  [5]string
	fmt.Println(a)

	var b [10]int
	fmt.Println(b)

	var c [15]bool
	fmt.Println(c)

	var d  [20]float64
	fmt.Println(d)

	var fillerArray  = [...]int{100,300,5:4999,233,499, 500}
	fmt.Println(fillerArray)

	var twoD = [2][3]int{
		{1,2,3},
		{1,2,3},
	}

	fmt.Println("2d: ", twoD)

	fmt.Println(twoD[0][0])
	fmt.Println(twoD[1][0])

}