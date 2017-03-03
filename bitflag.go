// Copyright 2017 Cosmin Albulescu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package bitflag

type Flag byte

// Set flag : Set(&flag, FLAG_B, FLAG_C)
func Set(f *Flag, opts ...Flag) {
	for _, o := range opts {
		*f |= o
	}
}

// Unset flag : Unset(&flag, FLAG_B, FLAG_C)
func Unset(f *Flag, opts ...Flag) {
	for _, o := range opts {
		*f ^= o
	}
}

// Flag is setted : Isset(&flag, FLAG_A), Isset(&flag, FLAG_A, FLAG_B), Isset(&flag, FLAG_A | FLAG_B)
func Isset(f Flag, opts ...Flag) (isset bool) {

	for _, o := range opts {
		if f&o == 0 {
			return false
		}
	}

	return true
}

// One of opts is setted in flag : OneOf(flag, FLAG_A, FLAG_B)
func One(f Flag, opts ...Flag) (isset bool) {

	for _, o := range opts {
		if f&o > 0 {
			return true
		}
	}

	return
}

func Clear(f *Flag) {
	*f = 0
}
