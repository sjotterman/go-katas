package kata

import "testing"

func TestPalindrome(t *testing.T) {
	testString := "a"
	want := true
	got := IsPalindrome(testString)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
