package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Address struct {
	City  string
	State string
}

type Employee struct {
	Name    string
	Age     int
	Address Address
}

type Student struct {
	Person // Embedded struct
	Grade  string
	School string
}

func main() {

	/* using struct literal */
	p2 := Person{"Bob", 30}
	fmt.Println(p2)

	/* using named fields */
	p3 := Person{Name: "Charlie", Age: 35}
	fmt.Println(p3)

	/* using new operator */
	p4 := new(Person)
	fmt.Println(*p4) // a zero-initialized struct
	p4.Name = "David"
	p4.Age = 40
	fmt.Println(*p4)

	/* anonymous structs */
	student := struct {
		Name  string
		Grade int
	}{"Eve", 90}

	fmt.Println(student)

	/* nested structs */
	/* we can omit fields too in named structs */
	emp := Employee{
		Name: "John Doe",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	fmt.Println(emp.Name, "lives in", emp.Address.City, emp.Address.State)

	/* embedded structs */
	s := Student{
		Person: Person{Name: "Alice", Age: 16},
		Grade:  "10th",
	}

	fmt.Println(s.Name, "is", s.Age, "years old and studies in", s.Grade)
}
