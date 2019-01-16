package bitflag

import "testing"

const (
	FLAG_A = Flag(1 << iota)
	FLAG_B
	FLAG_C
)

func TestSet(t *testing.T) {

	var flag Flag

	flag.Set(FLAG_A)
	flag.Set(FLAG_B)

	if !flag.IsSet(FLAG_A) {
		t.Fail()
	}

	if !flag.IsSet(FLAG_B) {
		t.Fail()
	}

	flag.Set(FLAG_A, FLAG_B)

	if !flag.IsSet(FLAG_A, FLAG_B) {
		t.Fail()
	}

	if !flag.OneOf(FLAG_B) {
		t.Fail()
	}
}

func TestUnset(t *testing.T) {

	var flag Flag

	flag.Set(FLAG_A, FLAG_B, FLAG_C)

	if !flag.IsSet(FLAG_A, FLAG_B, FLAG_C) {
		t.Fail()
	}

	flag.Unset(FLAG_B)

	if flag.IsSet(FLAG_B) {
		t.Fail()
	}

	if !flag.IsSet(FLAG_A, FLAG_C) {
		t.Fail()
	}

}

func TestClear(t *testing.T) {

	var flag Flag

	flag.Set(FLAG_A, FLAG_B, FLAG_C)

	if !flag.IsSet(FLAG_A, FLAG_B, FLAG_C) {
		t.Fail()
	}

	flag.Clear()

	if flag.OneOf(FLAG_A, FLAG_B, FLAG_C) {
		t.Fail()
	}
}
func TestSetAll(t *testing.T) {

	var flag Flag

	flag.SetAll()

	if !flag.IsSet(FLAG_A, FLAG_B, FLAG_C) {
		t.Fail()
	}
}
