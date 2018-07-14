package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func (p person) print() {
	fmt.Println(p)
}

func (p person) updateName(name string) {
	p.firstname = name
}

func (p *person) updateNameByPointer(name string) {
	(*p).firstname = name
}

func main() {

	enrique := person{
		firstname: "Enrique",
		lastname:  "Pereira",
		contactInfo: contactInfo{
			email:   "enrique.pereira.ssp@gmail.com",
			zipcode: 123456,
		},
	}

	enrique.print()
	fmt.Println("----- Do not change the value. Default behavior is 'pass by value'-----")
	enrique.updateName("HENRIQUE")
	enrique.print()
	fmt.Println("----- Changing value passing pointer -----")
	enriquePointer := &enrique
	enriquePointer.updateNameByPointer("HENRIQUE")
	enrique.print()

	fmt.Println("----- (Pointer shortcut) Changing value again passing value that is going to be converted to a pointer -----")
	enrique.updateNameByPointer("enrique")
	enrique.print()


}
