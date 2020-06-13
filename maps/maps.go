package main

import "fmt"

func main() {
	// To create an empty map, use the builtin `make`: make(map[key-type]value-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// Get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin `len` returns the number of key/value pairs
	fmt.Println("len:", len(m))

	// The builtin `delete` removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates
	// if the key was present in the map. This can be used to disambiguate between
	// missing keys and keys with zero values like 0 or "". Here we didn't need the
	// value itself, so we ignored it with the blank identifier `_`.
	fir, prs := m["k2"]
	fmt.Println("fir:", fir)
	fmt.Println("prs:", prs)

	// You can also declare and initialise a new map in the same line with this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
