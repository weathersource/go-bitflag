package bitflag

import "testing"

const FLAG_A uint64 = 1 << 1
const FLAG_B uint64 = 1 << 2
const FLAG_C uint64 = 1 << 3
const FLAG_D uint64 = 1 << 4

func TestExample(t *testing.T) {

	var flag Flag

	Set(&flag, FLAG_A)
	Set(&flag, FLAG_B, FLAG_C)
	Set(&flag, FLAG_C|FLAG_D)

	Clear(&flag)

	Set(&flag, FLAG_A, FLAG_B, FLAG_C, FLAG_D)

	Unset(&flag, FLAG_A)

	Unset(&flag, FLAG_B, FLAG_C)

	Unset(&flag, FLAG_A|FLAG_C)

	if Isset(flag, FLAG_A) {
		t.Fatal("A")
	}

	if Isset(flag, FLAG_B) {
		t.Fatal("B")
	}

	if Isset(flag, FLAG_C) {
		t.Fatal("C")
	}

	if !Isset(flag, FLAG_D) {
		t.Fatal("D")
	}
}

func TestSet(t *testing.T) {

	var flag Flag

	Set(&flag, FLAG_A)
	Set(&flag, FLAG_B)

	if !Isset(flag, FLAG_A) {
		t.Fail()
	}

	if !Isset(flag, FLAG_B) {
		t.Fail()
	}

	Set(&flag, FLAG_A, FLAG_B)

	if !Isset(flag, FLAG_A, FLAG_B) {
		t.Fail()
	}

	if !One(flag, FLAG_B) {
		t.Fail()
	}
}

func TestUnset(t *testing.T) {

	var flag Flag

	Set(&flag, FLAG_A, FLAG_B, FLAG_C)

	if !Isset(flag, FLAG_A, FLAG_B, FLAG_C) {
		t.Fail()
	}

	Unset(&flag, FLAG_B)

	if Isset(flag, FLAG_B) {
		t.Fail()
	}

	if !Isset(flag, FLAG_A, FLAG_C) {
		t.Fail()
	}

}

func TestClear(t *testing.T) {

	var flag Flag

	Set(&flag, FLAG_A, FLAG_B, FLAG_C)

	if !Isset(flag, FLAG_A, FLAG_B, FLAG_C) {
		t.Fail()
	}

	Clear(&flag)

	if One(flag, FLAG_A, FLAG_B, FLAG_C) {
		t.Fail()
	}
}
