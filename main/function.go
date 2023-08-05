package main

import "fmt"

func sayHello(s string) {
	fmt.Println("Hello " + s)
}

func sayBye(s string) {
	fmt.Println("Bye " + s)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	names := []string{"Mario", "luigi", "brawe"}
	cycleNames(names, sayHello)
	cycleNames(names, sayBye)

	ret := retunVal(names[0])
	fmt.Println(ret)
}

func retunVal(n string) string {
	return n
}
