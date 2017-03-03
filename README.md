# GoLang BitFlag

## Installation

`go get github.com/mvpninjas/go-bitflag`

## API

| Description        | Function  |
| ------------------ |:----------|
| Set flags | `bitflag.Set(f *Flag, opts ...Flag)` |
| Unset flags | `bitflag.Unset(f *Flag, opts ...Flag)` |
| All flags are setted | `bitflag.Isset(f Flag, opts ...Flag) bool` |
| One of flags is setted | `bitflag.One(f Flag, opts ...Flag) bool` |
| Clear all flags | `bitflag.Clear(f *Flag)` |


## Example

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
	
	bitflag.Set(&flag, C)
	
	if !bitflag.One(flag, D) {
		fmt.Println("E")
	}
}

```


## Example without this library 

```

type Flag byte

const (
	A = Flag(1 << iota)
	B
	B
)

var flag Flag

// set A
flag |= A

// check if A setted
if flag & A == 0 {
	// Flag setted
}

// remove A
flag ^= A

```
