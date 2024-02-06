package valid_anagram

import "testing"

func TestMyFunction(t *testing.T) {
	s := "anagram"
	n := "nagaram"
	result := isAnagram(s, n)
	expect := true
	if result != expect {
		t.Error()
	}
}
