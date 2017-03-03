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
	FLAG_A bitflag.Flag = 1 << bitflag.Flag(iota)
	FLAG_B
	FLAG_C
	FLAG_D
)

func main() {

	var flag bitflag.Flag

	bitflag.Set(&flag, FLAG_A)
	bitflag.Set(&flag, FLAG_B, FLAG_C)
	bitflag.Set(&flag, FLAG_C|FLAG_D)

	bitflag.Clear(&flag)

	bitflag.Set(&flag, FLAG_A, FLAG_B, FLAG_C, FLAG_D)

	bitflag.Unset(&flag, FLAG_A)

	bitflag.Unset(&flag, FLAG_B, FLAG_C)

	bitflag.Unset(&flag, FLAG_A|FLAG_C)

	if bitflag.Isset(flag, FLAG_A) {
		fmt.Println("A")
	}

	if bitflag.Isset(flag, FLAG_B) {
		fmt.Println("B")
	}

	if bitflag.Isset(flag, FLAG_C) {
		fmt.Println("C")
	}

	if !bitflag.Isset(flag, FLAG_D) {
		fmt.Println("D")
	}
}

```
