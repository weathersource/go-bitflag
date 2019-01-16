package bitflag_test

import (
	"fmt"

	bitflag "github.com/weathersource/go-bitflag"
)

func Example() {
	const (
		A = bitflag.Flag(1 << iota)
		B
		C
	)

	flags := A | C

	if flags.IsSet(A) {
		fmt.Println("A")
	}

	if flags.IsSet(B) {
		fmt.Println("B")
	}

	if flags.IsSet(C) {
		fmt.Println("C")
	}

	// Output: A
	// C
}

func ExampleSet() {
	const (
		A = bitflag.Flag(1 << iota)
		B
		C
	)

	flags := A

	flags.Set(C)

	if flags.IsSet(A, C) {
		fmt.Println("Success")
	}

	// Output: Success
}

func ExampleUnset() {
	const (
		A = bitflag.Flag(1 << iota)
		B
		C
	)

	flags := A

	flags.Unset(A)

	if !flags.IsSet(A) {
		fmt.Println("Success")
	}

	// Output: Success
}

func ExampleClear() {
	const (
		A = bitflag.Flag(1 << iota)
		B
		C
	)

	flags := A | C

	flags.Clear()

	if !flags.OneOf(A, C) {
		fmt.Println("Success")
	}

	// Output: Success
}
