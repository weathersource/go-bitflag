# GoLang BitFlag

## Installation

`go get github.com/albulescu/go-bitflag`

## Description

```

const FLAG_A uint = 1 << 1
const FLAG_B uint = 1 << 2
const FLAG_C uint = 1 << 3
const FLAG_D uint = 1 << 4

func TestExample(t *testing.T) {

	var flag uint64

	bitflag.Set(&flag, FLAG_A)
	bitflag.Set(&flag, FLAG_B, FLAG_C)
	bitflag.Set(&flag, FLAG_C|FLAG_D)

	bitflag.Clear(&flag)

	bitflag.Set(&flag, FLAG_A, FLAG_B, FLAG_C, FLAG_D)

	bitflag.Unset(&flag, FLAG_A)

	bitflag.Unset(&flag, FLAG_B, FLAG_C)

	bitflag.Unset(&flag, FLAG_A|FLAG_C)

	if bitflag.Isset(flag, FLAG_A) {
		t.Fatal("A")
	}

	if bitflag.Isset(flag, FLAG_B) {
		t.Fatal("B")
	}

	if bitflag.Isset(flag, FLAG_C) {
		t.Fatal("C")
	}

	if !bitflag.Isset(flag, FLAG_D) {
		t.Fatal("D")
	}
}

```
