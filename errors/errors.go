package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have the type `error`, a built-in interface
// `errors.New` constructs a basic `error` value with the given error message
// A `nil` value in the error position indicates that there was no error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// It's possible to use custom types as `errors` by implementing the Error() method on them.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// In this case we use &argError syntax to build a new struct, supplying values for the two fields
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// If you want to programmatically use the data in a custom error, you'll need to get
	// the error as an instance of the custom error type via `type assertion`.
	// A `type assertion` provides access to an interface value's underlying concrete value.
	// To test whether an interface value holds a specific type, a `type assertion` can return
	// two values: the underlying value and a boolean value that reports whether the assertion
	// succeeded. For example, `t, ok := i.(T)`
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
