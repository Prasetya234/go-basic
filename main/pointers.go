package main

import "fmt"

// pointer, menyebabkan pengantian langsung dengan alamat lokasi memori yang sama
// sehingga mengubah variable name tanpa menggunakan return
func updateName(x *string) {
	*x = "luigi"
}

func main() {
	name := "mario"

	m := &name

	fmt.Println("memory address ", m)
	fmt.Println("vault at memory address ", *m)

	fmt.Println(name)
	// merubah isi variable name dari function updateName berdasarkan lokasi memory
	updateName(m)
	fmt.Println(name)
}
