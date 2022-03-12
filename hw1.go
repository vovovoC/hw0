// Homework 1: Finger Exercises
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	//ParsePhone function
	fmt.Println("Testing ParsePhone function")
	fmt.Println("123-456-7890 -> ", ParsePhone("123-456-7890"))
	fmt.Println("1 2 3 4 5 6 7 8 9 0 -> ", ParsePhone("1 2 3 4 5 6 7 8 9 0"))

	//Anagram function
	fmt.Println("\nAnagram function")
	fmt.Println("golaNg & laGong are anagrams of each other:  ", Anagram("golaNg", "laGong"))
	fmt.Println("Lab1 & 2lba are anagrams of each other: ", Anagram("Lab1", "2lba"))

	//FindeEvens function
	fmt.Println("\nFindeEvens function")
	arr := []int{1, 2, 3, 4}
	fmt.Println("Find events from {1, 2, 3, 4}: ", FindEvens(arr))

	//SliceProduct function
	slice := []int{1, 2, 3}
	fmt.Println("\nSliceProduct function")
	fmt.Println("Sum of the {1, 2, 3}: ", SliceProduct(slice))

	//Unique function
	test := []int{1, 2, 2, 2, 3, 4, 4, 5, 6, 7, 7}
	fmt.Println("\nTesting Unique function")
	fmt.Println("Distinct elements from {1, 2, 2, 2, 3, 4, 4, 5, 6, 7, 7}: ", Unique(test))

	//InvertMap function
	n := map[string]int{"go": 1, "java": 2, "c++": 3}
	fmt.Println("\nTesting InvertMap function")
	fmt.Println("map[go:1 java:2 c++:3] -> ", InvertMap(n))

	//TopCharacters function
	fmt.Println("\nTesting TopCharacters function")
	char := "@@@hhhhh>>>"
	v := TopCharacters(char, 3)
	fmt.Println("Top char from @@@hh>>>, chars wich are >= 3: ")
	for key, value := range v {
		fmt.Println(string(key), ":", value)
	}
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	var res string = "("
	var index int = 1

	for i := 0; i < len(phone); i++ {
		if index == 4 {
			res = res + ")"
			index++
		} else if index == 5 {
			res = res + " "
			index++
		} else if index == 9 {
			res = res + "-"
			index++
		}
		if phone[i] >= '0' && phone[i] <= '9' {
			res = res + string(phone[i])
			index++
		}
	}
	return res
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	string1 := strings.Split(strings.ToLower(s1), "")
	sort.Strings(string1)

	string2 := strings.Split(strings.ToLower(s2), "")
	sort.Strings(string2)

	if len(string1) != len(string2) {
		return false
	} else {
		for i := 0; i < len(string1); i++ {
			if string1[i] != string2[i] {
				return false
			}
		}
	}
	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	for i := 0; i < len(e); i++ {
		if e[i]%2 != 0 {
			e = append(e[:i], e[i+1:]...)
		}
	}
	return e
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	var res int
	for i := 0; i < len(e); i++ {
		res += e[i]
	}
	return res
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	var check int = 0

	for i := 1; i < len(e); i++ {

		if e[check] == e[i] {
			e = append(e[:i], e[i+1:]...)
			i--
		} else {
			check++
		}

	}

	return e
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	vk := make(map[int]string)
	for key, value := range kv {
		vk[value] = key
	}
	return vk
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	var occurrences int = 0
	var j int = 0
	m := make(map[rune]int)

	str := strings.Split(s, "")
	sort.Strings(str)

	chars := []rune(strings.Join(str, ""))

	for i := 0; i < len(chars); i++ {
		if chars[j] == chars[i] {
			occurrences++

		} else {
			if occurrences >= k {
				m[rune(chars[j])] = occurrences
			}
			occurrences = 1
			j = i

		}

	}
	m[rune(chars[j])] = occurrences

	return m
}
