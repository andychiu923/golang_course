package main

import "fmt"


var g_groceries[]string

func add_grocery(a string) {
	fmt.Println("Capacity", cap(g_groceries))
	g_groceries=append(g_groceries, a)
}

func list_groceries() {
	fmt.Println("Grocery list is as follows: ")
	for i:=0; i<len(g_groceries); i++ {
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
	add_grocery("Holiday Ice Cream")
	list_groceries()
}
