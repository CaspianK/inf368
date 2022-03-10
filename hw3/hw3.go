// Homework 3: Interfaces
// Due February 14, 2017 at 11:59pm
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Feel free to use the main function for testing your functions
	// hello := map[string]string{
	// 	"hello":   "world",
	// 	"hola":    "mundo",
	// 	"bonjour": "monde",
	// }
	// for k, v := range hello {
	// 	fmt.Printf("%s, %s\n", strings.Title(k), v)
	// }

	p1 := NewPerson("Madi", "Hasen")
	p2 := NewPerson("Dana", "Aybar")
	personSlice := PersonSlice{p1, p2}
	for _, r := range personSlice {
		fmt.Println(r)
	}
	sort.Sort(personSlice)
	for _, r := range personSlice {
		fmt.Println(r)
	}

	sortString := SortString("kazak")
	fmt.Println(IsPalindrome(sortString))

	fmt.Println(Fold([]int{1, 2}, 0, f))
}

// Problem 1: Sorting Names
// Sorting in Go is done through interfaces!
// To sort a collection (such as a slice), the type must satisfy sort.Interface,
// which requires 3 methods: Len() int, Less(i, j int) bool, and Swap(i, j int).
// To actually sort a slice, you need to first implement all 3 methods on a
// custom type, and then call sort.Sort on your the PersonSlice type.
// See the Go documentation: https://golang.org/pkg/sort/ for full details.

// Person stores a simple profile. These should be sorted by alphabetical order
// by last name, followed by the first name, followed by the ID. You can assume
// the ID will be unique, but the names need not be unique.
// Sorting should be case-sensitive and UTF-8 aware.
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type PersonSlice []*Person

var id int = 1

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.
func NewPerson(first, last string) *Person {
	person := new(Person)
	person.ID = id
	person.FirstName = first
	person.LastName = last
	id++
	return person
}

func (p PersonSlice) Len() int {
	return len(p)
}

func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PersonSlice) Less(i, j int) bool {
	if p[i].LastName < p[j].LastName {
		return true
	}
	if p[i].LastName > p[j].LastName {
		return false
	}
	if p[i].FirstName < p[j].FirstName {
		return true
	}
	if p[i].FirstName > p[j].FirstName {
		return false
	}
	return p[i].ID < p[j].ID
}

// Problem 2: IsPalindrome Redux
// Using a function that simply requires sort.Interface, you should be able to
// check if a sequence is a palindrome. You may use, adapt, or modify your code
// from HW0. Note that the input does not need to be a string: any type which
// satisfies sort.Interface can (and will) be used to test. This means that the
// only functionality you should use should come from the sort.Interface methods
// Ex: [1, 2, 1] => true

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
type SortString string

func (s SortString) Len() int {
	return len([]rune(s))
}

func (s SortString) Swap(i, j int) {
	var t []rune = []rune(s)
	t[i], t[j] = t[j], t[i]
}

func (s SortString) Less(i, j int) bool {
	var t []rune = []rune(s)
	return t[i] < t[j]
}
func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		i++
		j--
	}
	return true
}

// Problem 3: Functional Programming
// Write a function Fold which applies a function repeatedly on a slice,
// producing a single value via repeated application of an input function.
// The behavior of Fold should be as follows:
//   - When s is empty, return v (default value)
//   - When s has 1 value (x0), apply f once: f(v, x0)
//   - When s has 2 values (x0, x1), apply f twice, from left to right: f(f(v, x0), x1)
//   - Continue this pattern recursively to obtain the final result.

// Fold applies a left to right application of f on s starting with v.
// Note the argument signature of f - func(int, int) int.
// This means f is a function which has 2 int arguments and returns an int.
func Fold(s []int, v int, f func(int, int) int) int {
	result := v
	if len(s) == 0 {
		return result
	}
	for _, r := range s {
		result = f(result, r)
	}
	return result
}

func f(i, j int) int {
	return i + j
}
