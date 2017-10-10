package passgen

import "crypto/rand"

// RandomString - generate char sequence from random elements letters
func RandomString(n int, letters []byte) string {
	bytes := make([]byte, n)
	rand.Read(bytes) // rand.Read always return (n, nil)

	l := byte(len(letters))
	for i, c := range bytes {
		bytes[i] = letters[c%l]
	}
	return string(bytes)
}
