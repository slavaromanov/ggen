package passgen

import "regexp"

var (
	empty []byte
	lower = []byte{
		97, 98, 99, 100, 101, 102, 103, 104, 105,
		106, 107, 108, 109, 110, 111, 112, 113, 114,
		115, 116, 117, 118, 119, 120, 121, 122,
	}
	upper = []byte{
		65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
		75, 76, 77, 78, 79, 80, 81, 82, 83, 84,
		85, 86, 87, 88, 89, 90,
	}
	puncts = []byte{
		33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
		44, 45, 46, 47, 58, 59, 60, 61, 62, 63, 64,
		91, 92, 93, 94, 95, 96, 123, 124, 125, 126,
	}
)

func newAlphabet(l, u, p bool) []byte {
	chars := []byte{}
	if l {
		chars = append(chars, lower...)
	}
	if u {
		chars = append(chars, upper...)
	}
	if p {
		chars = append(chars, puncts...)
	}
	return chars
}

// GetAlpha - decorator for generate alphabet
func GetAlpha(l, u, p bool, ex string) []byte {
	return regexp.MustCompile("["+ex+"]").ReplaceAll(newAlphabet(l, u, p), empty)
}
