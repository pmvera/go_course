package main

import "fmt"

func main() {
	// Not (!)
	fmt.Println(!true)
	fmt.Println(!false)

	// And (&&)
	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(false && false)

	// Or (||)
	fmt.Println(true || true)
	fmt.Println(false || true)
	fmt.Println(false || false)

	b := 2

	r := b == 2 && 4 > b
	fmt.Println("result:", r)

	r = b == 2 && !(4 > b)
	fmt.Println("result:", r)

	r = b == 2 || !(4 > b)
	fmt.Println("result:", r)

}
