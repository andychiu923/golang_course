package main

import "fmt"

const g_cap int = 5 	// total capacity of our list
var g_groceries[g_cap]string
var g_len int = 0		// total number of grocery item in our current array

func add_grocery(a string) {
	if g_len<g_cap {
		g_groceries[g_len] = a
		g_len++
	}else {
		fmt.Println("Too may items, now we are done for!!!!!")
	}
}

func list_groceried() {
	fmt.Println("Grocery list is as follows: ")
	for i:=0; i<g_len; i++ {
		fmt.Println(g_groceries[i])
	}
}

func main() {
	add_grocery("Bread")
	add_grocery("Cucumbers")
	add_grocery("Salt and Vinegar Potato Chips")
	add_grocery("Fruit Cake")
	add_grocery("Pokemon Game")
	add_grocery("Ice Cream")
	list_groceried()
}
