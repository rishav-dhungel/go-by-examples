package main

import "fmt"

func vals() (int, int){
	return 2,3
}

func category() (string, string, string){
	return "Action", "Adventure", "Fantasy"
}


func date() (int, string, int){
	return 2025, "August", 30
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	cat1, cat2, cat3 := category()
	fmt.Println(cat1)
	fmt.Println(cat2)
	fmt.Println(cat3)

	fmt.Println("=======================")
	_, month, _ := date()
	year, _, _ := date()
	_, _, day := date()

	fmt.Println(month)
	fmt.Println(year)
	fmt.Println(day)
}