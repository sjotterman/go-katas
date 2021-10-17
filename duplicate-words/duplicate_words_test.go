package kata

import "testing"

func TestFizzBuzz(t *testing.T) {

	type test struct {
		name  string
		input string
		want  string
	}

	tests := []test{
		{name: "empty string", input: "", want: ""},
		{name: "No duplicates", input: "one two three", want: "one two three"},
		{name: "Duplicate words", input: "alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta", want: "alpha beta gamma delta"},
		{name: "More duplicate words", input: "my cat is my cat fat", want: "my cat is fat"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RemoveDuplicateWords(test.input)
			if got != test.want {
				t.Errorf("For value %q: got %q, want %q", test.input, got, test.want)
			}
		})
	}
}
