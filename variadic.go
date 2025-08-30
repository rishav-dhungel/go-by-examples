package main

import "fmt"

func sum(nums ...int){
	fmt.Println(nums)
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


func mult(nums ...int){
	total := 1

	for _, num := range nums{
		total *= num
	}

	fmt.Println(total)

}

func filter5(nums ...int) []int{
	var newList []int
	for _, num := range nums{
		if num >5{
			newList = append(newList, num)
		}
	}

	return newList
}


func main() {
	sum(1,2)
	sum(1,2,3,4)

	mult(1,2,3)
	newList := filter5(12,3,2,41,345,345,342,3,234,2342)
	fmt.Println(newList)
}