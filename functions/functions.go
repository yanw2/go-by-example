package main

import "fmt"

// Go requires explicit returns
func plus(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type,
// you might omit the type name for the like-typed parameters up to
// the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Go has built-in support for multiple return values
// The (int, int) in this function signature shows that the function
// returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

// Variadic functions ca be called with any number of trailing arguments.
// Here's a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Go supports anonymous functions, which can form closures. Anonymous functions are
// useful when you want to define a function inline without having to name it.
// This function `intSeq` returns another function, which we define anonymously in the
// body of `intSeq`. The returned function closes over the variable `i` to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Go supports recursive functions. Here's a classic factorial example.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	res1 := plus(1, 2)
	fmt.Println("1+2 =", res1)

	res2 := plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res2)

	// Here we use the 2 different return values from the call with multiple assignment
	// If you only want a subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)

	// Variadic functions
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// We call `intSeq`, assigning the result (a function) to `nextInt`. This function
	// value captures its own `i` value, which will be updated each time we call `nextInt`.
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())

	// Here's a classic factorial example
	fmt.Println(fact(7))
}
