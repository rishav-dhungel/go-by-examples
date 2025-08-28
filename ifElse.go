package main

import "fmt"

func main() {

	if 17 % 3 == 0{
		fmt.Println("17 is divisible by 3. YAY!")
	}else{
		fmt.Println("17 is not divisible by 3. NAY!")
	}

	var a int  = 90

	if a == 10 || a == 100{
		fmt.Println("well the answer is either a equals to 10 or a equals to 100.")
	}else if a > 100 {
		fmt.Println("neither a equals to 10 or 100 but is greater than 100")
	}else{
		fmt.Println("We have no interest in knowing what a is at this point.")
	}
}