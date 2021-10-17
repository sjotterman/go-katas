package kata

import "testing"

func TestPalindrome(t *testing.T) {

	type test struct {
		name  string
		input string
		want  string
	}

	tests := []test{
		{name: "one letter", input: "a", want: "a"},
		{name: "two matching letters", input: "aa", want: "aa"},
		{name: "two different letters", input: "ab", want: "ba"},
		{name: "word", input: "world", want: "dlrow"},
		{name: "unicode", input: "Привет", want: "тевирП"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ReverseString(test.input)
			if got != test.want {
				t.Errorf("For string %q: got %q, want %q", test.input, got, test.want)
			}
		})
	}
}
