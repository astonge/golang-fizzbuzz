package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(fb int) {
	slices := make([]string, fb+1)
	for i := 1; i < len(slices); i++ {
		if (i%3 == 0) && (i%5 == 0) {
			slices[i] = "FizzBuzz"
		} else if (i%3 == 0) && !(slices[i] == "FizzBuzz") {
			slices[i] = "Fizz"
		} else if (i%5 == 0) && !(slices[i] == "FizzBuzz") {
			slices[i] = "Buzz"
		} else {
			slices[i] = strconv.Itoa(i)
		}
	}
	fmt.Println(slices)
}

func main() {
	fizzbuzz(3)
	fizzbuzz(16)
	fizzbuzz(100)
}