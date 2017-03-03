package bitflag

// Set flag : Set(&flag, FLAG_B, FLAG_C)
func Set(f *uint64, opts ...uint64) {
	for _, o := range opts {
		*f |= (1 << o)
	}
}

// Unset flag : Unset(&flag, FLAG_B, FLAG_C)
func Unset(f *uint64, opts ...uint64) {
	for _, o := range opts {
		*f &= uint64(^(1 << o))
	}
}

// Flag is setted : Isset(&flag, FLAG_A), Isset(&flag, FLAG_A, FLAG_B), Isset(&flag, FLAG_A | FLAG_B)
func Isset(f uint64, opts ...uint64) (isset bool) {

	for _, o := range opts {
		if (f & (1 << o)) == 0 {
			return false
		}
	}

	return true
}

// One of opts is setted in flag : OneOf(flag, FLAG_A, FLAG_B)
func One(f uint64, opts ...uint64) (isset bool) {

	for _, o := range opts {
		if (f & (1 << o)) > 0 {
			return true
		}
	}

	return
}

func Clear(f *uint64) {
	*f = 0
}
