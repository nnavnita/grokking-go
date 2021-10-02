package main

import "fmt"

func getFactorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * getFactorial(num-1)
	}
}

func main() {
	fmt.Printf("Factorial of %d is %d\n", 5, getFactorial(5))
	fmt.Printf("Factorial of %d is %d\n", 10, getFactorial(10))
}
