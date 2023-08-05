package main

import (
	"fmt"
	"strings"
)

func getInitials(fullname string) (string, string) {
	n := strings.ToUpper(fullname)
	names := strings.Split(n, " ")

	var initalize []string
	for _, v := range names {
		initalize = append(initalize, v)
	}

	if len(initalize) <= 1 {
		return initalize[0], "-"
	}
	return initalize[0], initalize[1]
}

func main() {
	firstName, lastName := getInitials("Jhon Doe")
	fmt.Println(firstName, lastName)

	firstName1, lastName2 := getInitials("Mario")
	fmt.Println(firstName1, lastName2)
}
