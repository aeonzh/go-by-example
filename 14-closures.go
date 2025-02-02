package main

import "fmt"

// from Wikipedia: a closure allows the function to access those _captured variables_
// through the closure's copies of their values or references,
// even when the function is invoked outside their scope
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := intSeq()

	// Here we use the captured i from nextInt()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Here we use the captured i from newInts()
	newInts := intSeq()
	fmt.Println(newInts())
}
