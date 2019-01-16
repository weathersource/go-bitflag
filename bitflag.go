package bitflag

// Flag is a container for bit flags.
type Flag uint64

// Set adds flags to Flag.
func (f *Flag) Set(flags ...Flag) {
	for _, o := range flags {
		*f |= o
	}
}

// Unset removes flags from Flag.
func (f *Flag) Unset(flags ...Flag) {
	for _, o := range flags {
		*f ^= o
	}
}

// IsSet returns true if all flags are set in Flag, false otherwise.
func (f Flag) IsSet(flags ...Flag) bool {
	for _, o := range flags {
		if f&o == 0 {
			return false
		}
	}
	return true
}

// OneOf returns true if any flag in flags is set in Flag, false otherwise.
func (f Flag) OneOf(flags ...Flag) bool {
	for _, o := range flags {
		if f&o > 0 {
			return true
		}
	}
	return false
}

// Clear removes all flags from Flag
func (f *Flag) Clear() {
	*f = 0
}

// SetAll sets all flags in Flag
func (f *Flag) SetAll() {
	f.Clear()
	*f -= 1
}
