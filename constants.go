package main

import(
	"fmt"
	"math"
)

const s string = "THIS IS CONSTANT"

func main() {
	fmt.Println(s)

	const age = 250000000
	fmt.Println(math.Sin(age))
	fmt.Println(int64(age))

	const sampleUnit = 2e10

	fmt.Println(sampleUnit)
}