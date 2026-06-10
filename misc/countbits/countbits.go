package countbits

// CountBits returns the number of set bits (1s) in the binary
// representation of n.
//
// Example:
//
//	CountBits(13) // 1101 → 3
func CountBits(n int) int {
	var counter int
	for n > 0 {
		n = n & (n - 1)
		counter++
	}
	return counter
}
