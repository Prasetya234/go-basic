package main

import "fmt"

func main() {
	object := map[string]any{
		"name":    "Carrolina",
		"age":     22,
		"address": "Bali",
	}

	object["name"] = "Jhon"

	fmt.Println(object)
}
