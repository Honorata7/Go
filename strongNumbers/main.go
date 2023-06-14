package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func getAsciiCodes(name string) []string {
	var asciiCodes []string
	for _, char := range name {
		asciiCodes = append(asciiCodes, strconv.Itoa(int(char)))
	}
	return asciiCodes
}

func factorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	return big.NewInt(n).Mul(factorial(n-1), big.NewInt(n))
}

func containsAllFactors(fact string, factors []string) bool {
	for _, f := range factors {
		if !containsFactor(fact, f) {
			return false
		}
	}
	return true
}

func containsFactor(str, factor string) bool {
	return len(str) >= len(factor) && (str[:len(factor)] == factor || containsFactor(str[1:], factor))
}

func getFirstNumberWithFactors(factors []string, max int64) (string, error) {
	for n := int64(1); n <= max; n++ {
		fact := factorial(n).String()
		if containsAllFactors(fact, factors) {
			return strconv.FormatInt(n, 10), nil
		}
	}
	return "", errors.New("no number found with all factors")
}

func getNumbersWithFactors(factors []string, max int64) []string {
	var numbers []string

	for n := int64(1); n <= max; n++ {
		fact := factorial(n).String()
		if containsAllFactors(fact, factors) {
			numbers = append(numbers, strconv.FormatInt(n, 10))
		}
	}

	return numbers
}

func fibonacci(n int, count map[int]int) int {
	if n <= 1 {
		return n
	}

	// Increment the count for the current Fibonacci function call
	count[n]++

	return fibonacci(n-1, count) + fibonacci(n-2, count)
}

func main() {

	name := "honora"
	asciiCodes := getAsciiCodes(name)

	fmt.Println("The ascii codes for nick are: ", asciiCodes)

	max := int64(1000)
	numbers := getNumbersWithFactors(asciiCodes, max)

	number := ""
	var err error
	if number, err = getFirstNumberWithFactors(asciiCodes, max); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The first number with a factorial that contains all factors is: %v\n", number)
	}

	if len(numbers) > 0 {
		fmt.Printf("The following numbers have factorials that contain all factors: %v\n", numbers)
	} else {
		fmt.Println("No numbers have factorials that contain all factors")
	}

	n := 30
	count := make(map[int]int)

	fibonacci30 := fibonacci(n, count)
	fmt.Printf("Fibonacci(%d): %d\n", n, fibonacci30)

	fmt.Println("Function call counts:")
	target, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	temp := 0
	closest := 0
	minDiff := math.Abs(float64(target - closest))
	for i, c := range count {
		fmt.Printf("Fibonacci(%d) called %d times\n", i, c)

		diff := math.Abs(float64(target - c))
		if diff < minDiff {
			minDiff = diff
			closest = c
			temp = i
		}
	}

	fmt.Println("Closest value to", target, "is", closest)
	fmt.Println("It means that the other number is", temp)

}
