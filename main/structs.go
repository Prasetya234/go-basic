package main

import "fmt"

type student struct {
	id          int
	email       string
	phoneNumber string
	age         int
	address     string
}

func Students(email string, phoneNumber string, age int) student {
	res := student{
		id:          1,
		email:       email,
		phoneNumber: phoneNumber,
		age:         age,
		address:     "-",
	}

	return res
}

func (s student) getEmail() string {
	return s.email
}

func (s *student) setEmail(email string) {
	s.email = email
}

func main() {
	res := Students("prasetya@gmail.com", "123983883", 18)
	fmt.Println(res)
	res.setEmail("pras@gmail.com")
	email := res.getEmail()
	fmt.Println(email)

}
