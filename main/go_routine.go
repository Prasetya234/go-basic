package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)       // jumlah go routine
	defer wg.Wait() // di kasih await di paling bawah dengan (defer)
	go asyncFunc(12, "Hello WOrld")
	go asyncFunc(5, "Hello Guys")
}

func asyncFunc(lim int, word string) {
	defer wg.Done() // setelah function go routine selesai harus di done
	for i := 0; i < lim; i++ {
		fmt.Println(word)
		time.Sleep(time.Millisecond * 100) // jeda 100 millisecond
	}
}

// go routine is function asynchronous
