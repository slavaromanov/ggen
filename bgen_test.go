package passgen

import "testing"

func TestRandomString(t *testing.T) {
	if l := len(RandomString(16, newAlphabet(true, true, true))); l != 16 {
		t.FailNow()
	}
}
