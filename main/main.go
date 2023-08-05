package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// function remove index
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	// vairable

	var name string = "Jhon" // menyebutkan tipe data
	var name2 = "Doe"        // tanpa menyebutkan tipe data
	var name3 string         // membuat variable dengan awalan kosong

	name4 := "hello" // membuat variable tanpa kata kunci

	fmt.Println("Variable")
	fmt.Println(name + name2)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)

	fmt.Println("operatior penjumlahan")
	number1 := 10
	number2 := "100"

	// konversi tipedata
	number3, _ := strconv.Atoi(number2)

	fmt.Println(number1 + number3)

	fmt.Println("Array")
	// array static
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr, len(arr))

	// dynamic array
	arr2 := []string{"Jhon", "Doe", "Mark"}
	arr2[1] = "Doee"

	// add array
	arr2 = append(arr2, "Nana")
	fmt.Println(arr2)

	// remove index by index
	arr2 = RemoveIndex(arr2, 0)
	fmt.Println(arr2)

	// get by index
	fmt.Println(arr2[0])

	// string function
	fmt.Println("String function")
	text := "hello everyone guys!"
	text = strings.ToUpper(text)                       // uppercase
	searchTextValid := strings.Contains(text, "hello") // serach text
	arrayText := strings.Split(text, " ")              // split
	fmt.Println(text)
	fmt.Println(searchTextValid)
	fmt.Println(arrayText)

	fmt.Println("Shorting array")
	ages := []int{19, 18, 21, 20, 22, 17}
	names := []string{"Jhon", "Elon", "Bill", "Mark"}

	sort.Ints(ages)     // asc
	sort.Strings(names) // asc
	fmt.Println(ages)
	fmt.Println(names)
}
