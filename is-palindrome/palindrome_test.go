package kata

import "testing"

func TestPalindrome(t *testing.T) {

	type test struct {
		name  string
		input string
		want  bool
	}

	tests := []test{
		{name: "one letter", input: "a", want: true},
		{name: "two matching letters", input: "aa", want: true},
		{name: "two different letters", input: "ab", want: false},
		{name: "palindrome with capital", input: "Abba", want: true},
		{name: "non-palindrome", input: "hello", want: false},
	}

	for _, test := range tests {
		got := IsPalindrome(test.input)
		if got != test.want {
			t.Errorf("%s: For string \"%q\": got %v, want %v", test.name, test.input, got, test.want)
		}
	}
}
