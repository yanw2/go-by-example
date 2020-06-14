package main

import "fmt"

// Go's structs are typed collections of fields. They're useful for grouping
// data together to form records.
type person struct {
	name string
	age  int
}

// Go supports methods defined on struct types
type rect struct {
	width, height int
}

// This area method has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// Create a new struct
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initialising a struct
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// Struct methods
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// Go automatically handles conversion between values and pointers for methods calls.
	// You may want to use a pointer receiver type to avoid copying on method calls or
	// to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
