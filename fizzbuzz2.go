package main

import (
	"fmt"
	"strconv"
)

func isFizzBuzz(value int) string {
	if (value%3 == 0) && (value%5 == 0) {
		return "FizzBuzz"
	}

	if (value%3 == 0) && !(strconv.Itoa(value) == "FizzBuzz") {
		return "Fizz"
	}

	if (value%5 == 0) && !(strconv.Itoa(value) == "FizzBuzz") {
		return "Buzz"
	}

	return strconv.Itoa(value)
}

func fizzbuzz(n int) []string {
	slices := make([]string, n+1)
	for i := 1; i < len(slices); i++ {
		slices[i] = isFizzBuzz(i)
	}
	return slices
}

func main() {
	fmt.Println(fizzbuzz(32))
}