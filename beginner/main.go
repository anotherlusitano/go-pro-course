package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Maria", Age: 40}
	fmt.Printf("Person: %+v\n", person)

	// Anonymous struct
	employee := struct {
		name string
		id   int
	}{
		name: "Bob",
		id:   65,
	}

	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact{
		Name: "Marcos",
		Address: Address{
			Street: "47 Main street",
			City:   "Anytown",
		},
	}

	fmt.Println(employee)
	fmt.Println(contact)

	fmt.Printf("Name before: %s\n", person.Name)

	person.modifyPersonName("John Cena")

	fmt.Printf("Name after: %s\n", person.Name)

	x := 20
	ptr := &x
	fmt.Printf("value of x: %d and address of x %p\n", x, ptr)
	*ptr = 30
	fmt.Printf("value of new x: %d and address of x %p\n", x, ptr)
}

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Printf("Inside scope new name: %s\n", p.Name)
}
