package kata

import "testing"

func TestFizzBuzz(t *testing.T) {

	type test struct {
		name  string
		input int
		want  string
	}

	tests := []test{
		{name: "1", input: 1, want: "1"},
		{name: "2", input: 2, want: "2"},
		{name: "3", input: 3, want: "Fizz"},
		{name: "4", input: 4, want: "4"},
		{name: "5", input: 5, want: "Buzz"},
		{name: "6", input: 6, want: "Fizz"},
		{name: "8", input: 8, want: "8"},
		{name: "9", input: 9, want: "Fizz"},
		{name: "10", input: 10, want: "Buzz"},
		{name: "12", input: 12, want: "Fizz"},
		{name: "15", input: 15, want: "FizzBuzz"},
		{name: "20", input: 20, want: "Buzz"},
		{name: "30", input: 30, want: "FizzBuzz"},
		{name: "33", input: 33, want: "Fizz"},
		{name: "45", input: 45, want: "FizzBuzz"},
		{name: "99", input: 99, want: "Fizz"},
		{name: "100", input: 100, want: "Buzz"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NthFizzBuzz(test.input)
			if got != test.want {
				t.Errorf("For value %d: got %v, want %q", test.input, got, test.want)
			}
		})
	}
}
