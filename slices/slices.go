package main

import "fmt"

func main() {
	// Unlike array, slices are typed only by the elements they contain.
	// To craete an emtpy slice with non-zero length, use the builtin `make`
	// Here we make a slice of `strings` of length 3 (initially zero-valued)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected
	fmt.Println("len:", len(s))

	// append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copy'd
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax `slice[low:high]`
	l1 := s[2:5]
	fmt.Println("sl1:", l1)

	// This slices up to (but excluding) s[5]
	l2 := s[:5]
	fmt.Println("sl2:", l2)

	// This slices up from (and including) s[2]
	l3 := s[2:]
	fmt.Println("sl3", l3)

	// We can declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
