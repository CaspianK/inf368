// Homework 2: Object Oriented Programming
// Due February 7, 2017 at 11:59pm
package main

import (
	"fmt"
	"log"
)

func main() {
	// Feel free to use the main function for testing your functions
	// world := struct {
	// 	English string
	// 	Spanish string
	// 	French  string
	// }{
	// 	"world",
	// 	"mundo",
	// 	"monde",
	// }
	// fmt.Printf("Hello, %s/%s/%s!", world.English, world.Spanish, world.French)
	var p Price = 2595
	s := p.String()
	fmt.Println(s)
	RegisterItem(Prices, "bread", Price(520))
	fmt.Println(Prices)
	cart := Cart{[]string{"milk", "bread", "eggs"}, Price(713)}
	fmt.Println(cart.hasMilk())
	fmt.Println(cart.HasItem("bread"))
	cart.AddItem("chocolate")
	fmt.Println(cart.TotalPrice)
	cart.Checkout()
	fmt.Println(cart)
}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	dollars := int(p) / 100
	cents := int(p) - dollars*100
	return fmt.Sprintf("$%d.%d", dollars, cents)
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	if _, contains := prices[item]; contains {
		log.Println("Error! The item is already in the map!")
	}
	prices[item] = price
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	for _, r := range c.Items {
		if r == "milk" {
			return true
		}
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, r := range c.Items {
		if r == item {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	if _, contains := Prices[item]; !contains {
		log.Println("Error! The item has not been found!")
	} else {
		c.Items = append(c.Items, item)
		c.TotalPrice += Prices[item]
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	fmt.Println("The total is", c.TotalPrice)
	c.Items = []string{}
	c.TotalPrice = 0
}
