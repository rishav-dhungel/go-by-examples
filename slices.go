package main

import(
	"fmt"
)

func main() {
	var s []string
	s = append(s, "We have added something to slice, I thought slice ment, slicing array in python but it ain't that.")
	s = append(s,"one", "two")
	fmt.Println(s)
	fmt.Println(len(s))


}