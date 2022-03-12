package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Feel free to use the main function for testing your functions
	world := struct {
		English string
		Spanish string
		French  string
	}{
		"world",
		"mundo",
		"monde",
	}
	fmt.Printf("Hello, %s/%s/%s!", world.English, world.Spanish, world.French)

	var price1 Price = 45576
	fmt.Println("\nTest String()", price1.String())

	var price Price = 219
	fmt.Println("\nTest RegisterItem(): ")
	RegisterItem(Prices, "eggs", price)

	var c Cart
	c.Items = []string{"milk", "coffee"}

	fmt.Println("\nTest hasMilk() in cart items {'milk', 'coffee'}")
	fmt.Println("Has milk: ", c.hasMilk())

	fmt.Println("\nTest hasItem() in cart items {'milk', 'coffee'}")
	fmt.Println("Has coffee: ", c.HasItem("coffee"))
	fmt.Println("Has bread: ", c.HasItem("bread"))

	fmt.Println("\nTest AddTem()\nbread add to card")
	c.AddItem("bread")
	fmt.Println("\nWater add to card: ")
	c.AddItem("water")
	fmt.Println(Prices)

	fmt.Println("\nTest Checkout()")
	c.Checkout()

}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	v := strconv.FormatInt(int64(p), 10)
	res := "$"
	for i := 0; i < len(v); i++ {
		if i == len(v)-2 {
			res = res + "." + string(v[i])
		} else {
			res = res + string(v[i])
		}

	}

	return res
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"coffee":        150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user

func RegisterItem(prices map[string]Price, item string, price Price) {
	if prices[item] == price {
		log.Println("Item is already in the prices map")

	}
	prices[item] = price
	fmt.Println(prices)

}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	return c.HasItem("milk")

}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, v := range c.Items {
		if v == item {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	if Prices[item] == 0 {
		log.Println(item, " is not found in the prices map")
	} else {
		c.Items = append(c.Items, item)
		fmt.Println(c.Items)

	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	for _, v := range c.Items {
		if Prices[v] != 0 {
			c.TotalPrice += Prices[v]
		}
	}
	fmt.Println("Cart:", c.Items, "\nFinal cart balance ", c.TotalPrice)
	c.TotalPrice = 0
	c.Items = nil

}
