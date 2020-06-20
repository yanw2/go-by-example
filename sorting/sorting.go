package main

import (
	"fmt"
	"sort"
)

// In order to sort by a custom function in Go, we need a corresponding type. Here we've created a `byLength`
// type that is just an alias for the builtin `[]string` type.
type byLength []string

// We implement `sort.Interface -len, Less and Swap` - on our type so we can use the `sort` package's generic
// Sort function. `Len` and `Swap` will usually be similar across types and `Less` will hold the actual custom
// sorting logic. In our case, we want to sort in order of increasing string length, so we use `len(s[i])` and
// `len(s[j])` here.
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// Built-in sorting
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)

	// Custom sorting
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	// By following this smae pattern of creating a custom type, implementing the three `Interface` methods on
	// that type, and then calling `sort.Sort` on a collection of that custom type, we can sort Go slices by
	// arbitrary functions.
}
