package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/// Read from argument
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	res := Sum(a, b)

	fmt.Printf("Sum of %d and %d is %d ", a, b, res)

}

func Sum(a, b int) int {
	return a + b
}
