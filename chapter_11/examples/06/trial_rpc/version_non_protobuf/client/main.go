package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"First", "A first item"}
	b := Item{"Second", "A second item"}
	c := Item{"Third", "A third item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Database: ", db) // at 1st call

	client.Call("API.EditItem", Item{"Second", "A new second item"}, &reply)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database: ", db) // at 2nd call

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first item: ", reply)

}

/* version 2
// Extracted constants
var items = []Item{
	{"First", "A first item"},
	{"Second", "A second item"},
	{"Third", "A third item"},
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error:", err)
	}

	// Add all items using a loop
	for _, item := range items {
		if err := client.Call("API.AddItem", item, &reply); err != nil {
			log.Println("AddItem error:", err)
		}
	}

	// Get DB after additions
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database (after add):", db)

	// Edit second item
	if err := client.Call("API.EditItem", Item{"Second", "A new second item"}, &reply); err != nil {
		log.Println("EditItem error:", err)
	}

	// Delete third item
	c := "Third"
	if err := client.Call("API.DeleteItem", c, &reply); err != nil {
		log.Println("DeleteItem error:", err)
	}

	// Get DB after edits and delete
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database (after edit/delete):", db)

	// Fetch by name
	if err := client.Call("API.GetByName", "First", &reply); err != nil {
		log.Println("GetByName error:", err)
	} else {
		fmt.Println("First item:", reply)
	}
}
*/
