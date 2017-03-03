package bitflag

import "testing"

const FLAG_A uint = 1 << 1
const FLAG_B uint = 1 << 2
const FLAG_C uint = 1 << 3

func TestSet(t *testing.T) {

	var flag uint

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

	var flag uint

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

	var flag uint

	Set(&flag, FLAG_A, FLAG_B, FLAG_C)

	if !Isset(flag, FLAG_A, FLAG_B, FLAG_C) {
		t.Fail()
	}

	Clear(&flag)

	if One(flag, FLAG_A, FLAG_B, FLAG_C) {
		t.Fail()
	}
}
