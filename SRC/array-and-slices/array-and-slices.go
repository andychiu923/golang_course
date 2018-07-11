package main

import "fmt"


var g_groceries[]string

func add_grocery(a ...string) {
	fmt.Println("Capacity", cap(g_groceries))
	for _,d:=range a {
		g_groceries=append(g_groceries,d)
	}
	//g_groceries=append(g_groceries, a)
}

func list_groceries() {
	fmt.Println("Grocery list is as follows: ")
	/*for i:=0; i<len(g_groceries); i++ {
		fmt.Println(g_groceries[i])
	}*/

	for _,d:=range g_groceries {
		fmt.Println(d)
	}
}

func main() {
	add_grocery("Bread","Cucumbers","Salt and Vinegar Potato Chips","Fruit Cake,Pokemon Game","Holiday Ice Cream", "Corn", "Peas")

	list_groceries()
}
