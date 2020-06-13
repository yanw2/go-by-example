package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("hello world" + ". hello golang.")

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	const n = 5000000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(math.Sin(n))

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
