package utils

import "math"

// AddU16 adds two uint16 together and tells if the operation overflowed the type
func AddU16(a, b uint16) (uint16, bool) {
	return a + b, math.MaxUint16-a < b || math.MaxUint16-b < a
}
