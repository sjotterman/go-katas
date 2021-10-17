package kata

import "testing"

func TestRemoveFirstLast(t *testing.T) {

	type test struct {
		input string
		want  string
	}

	tests := []test{
		{"eloquent", "loquen"},
		{"country", "ountr"},
		{"person", "erso"},
		{"place", "lac"},
		{"Sam", "a"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := RemoveFirstLast(test.input)
			if got != test.want {
				t.Errorf("For value %q: got %q, want %q", test.input, got, test.want)
			}
		})
	}
}
