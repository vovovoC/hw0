// Homework 0: Hello Go!
// Due January 24, 2017 at 11:59pm
package main

import "fmt"

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println(Fizzbuzz(1))  // ""
	fmt.Println(Fizzbuzz(15)) // "FizzBuzz"
	fmt.Println(Fizzbuzz(25)) // "Buzz"
	fmt.Println(Fizzbuzz(27)) // "Fizz"

	fmt.Println(IsPrime(27)) // false
	fmt.Println(IsPrime(23)) // true

	fmt.Println(IsPalindrome("hello")) // false
	fmt.Println(IsPalindrome("pooop")) // true
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	// TODO
	var ans = ""
	if n%3 == 0 {
		ans += "Fizz"
	}
	if n%5 == 0 {
		ans += "Buzz"
	}
	return ans
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	// TODO
	if n < 2 {
		return false
	}
	for t := 2; t < n; t++ {
		if n%t == 0 {
			return false
		}
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	// TODO
	for t := 0; t < len(s)/2; t++ {
		if s[t] != s[len(s)-1-t] {
			return false
		}
	}
	return true
}
