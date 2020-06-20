package main

import "os"

func main() {
	// A `panic` typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that
	// shouldn't occur during normal operation, or that we aren't prepared to handle gracefully.
	panic("a problem")

	// A common use of `panic` is to abort if a function returns an error valuethat we don't know how to
	// (or want to) handle. Here's an example of panicking if we get an unexpected error when creating a new file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
