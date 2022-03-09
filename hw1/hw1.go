// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println(ParsePhone("1 2 3 4 5 6 7 8 9 0"))
	fmt.Println(Anagram("New York Times", "monkeys write"))
	fmt.Println(FindEvens([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(SliceProduct([]int{1, 2, 3}))
	fmt.Println(Unique([]int{1, 1, 1, 2, 3, 4, 2, 2, 5, 6}))
	fmt.Println(InvertMap(map[string]int{
		"henry": 4,
		"sarah": 12,
		"jim":   5,
	}))
	fmt.Println(TopCharacters("hello there", 2))

}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	result := "("
	i := 0
	for _, r := range phone {
		if r >= 48 && r <= 57 {
			result += string(r)
			i++
			if i == 3 {
				result += ") "
			}
			if i == 6 {
				result += "-"
			}
		}
	}
	return result
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	var slice1 []int
	var slice2 []int
	for _, r := range strings.ToLower(s1) {
		if r != 32 {
			slice1 = append(slice1, int(r))
		}
	}
	for _, r := range strings.ToLower(s2) {
		if r != 32 {
			slice2 = append(slice2, int(r))
		}
	}
	sort.Ints(slice1)
	sort.Ints(slice2)
	return reflect.DeepEqual(slice1, slice2)
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var result []int
	for _, r := range e {
		if r%2 == 0 {
			result = append(result, r)
		}
	}
	return result
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	result := 1
	for _, r := range e {
		result *= r
	}
	return result
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	var result []int
	for _, r := range e {
		contains := false
		for _, r2 := range result {
			if r == r2 {
				contains = true
				break
			}
		}
		if !contains {
			result = append(result, r)
		}
	}
	return result
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	result := make(map[int]string)
	for k, v := range kv {
		result[v] = k
	}
	return result
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	result := make(map[rune]int)
	for _, r := range s {
		if r != 32 {
			if _, contains := result[r]; !contains {
				result[r] = 1
			} else {
				result[r]++
			}
		}
	}
	for key := range result {
		if result[key] < k {
			delete(result, key)
		}
	}
	return result
}
