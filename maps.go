package main

import(
	"fmt"
)

func main() {

	m := make(map[string]int)  // make(map[key-type)val-type
	second := make(map[int]string)
	m["One"] = 1
	second[1] = "One"

	fmt.Println(m["One"])
	fmt.Println(second[1])
	fmt.Println(m)
	fmt.Println(second)

	n := map[string]string{
		"one": "one",
		"bar":"bar",
		"Username":"admin",
	}

	fmt.Println("\rThe content of the map n is as follow\n",n)


}