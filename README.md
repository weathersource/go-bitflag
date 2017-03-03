# GoLang BitFlag

## Installation

`go get github.com/albulescu/go-bitflag`

## Description

```
package main

import (
	"fmt"
	bitflag "github.com/albulescu/go-bitflag"
)

const (
	A bitflag.Flag = 1 << bitflag.Flag(iota)
	B
	C
	D
)

func main() {

	var flag bitflag.Flag

	bitflag.Set(&flag, A)
	bitflag.Set(&flag, B, C)
	bitflag.Set(&flag, C|D)

	bitflag.Clear(&flag)

	bitflag.Set(&flag, A, B, C, D)

	bitflag.Unset(&flag, A)

	bitflag.Unset(&flag, B, C)

	bitflag.Unset(&flag, A|C)

	if bitflag.Isset(flag, A) {
		fmt.Println("A")
	}

	if bitflag.Isset(flag, B) {
		fmt.Println("B")
	}

	if bitflag.Isset(flag, C) {
		fmt.Println("C")
	}

	if !bitflag.Isset(flag, D) {
		fmt.Println("D")
	}
}

```
