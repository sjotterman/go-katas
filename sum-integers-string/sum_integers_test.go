package kata

import "testing"

func TestSumIntegers(t *testing.T) {

	type test struct {
		input string
		want  int
	}

	tests := []test{
		{"12.4", 16},
		{"h3ll0w0rld", 3},
		{"2 + 3 = ", 5},
		{"Our company made approximately 1 million in gross revenue last quarter.", 1},
		{"The Great Depression lasted from 1929 to 1939.", 3868},
		{"Dogs are our best friends.", 0},
		{"The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", 3635},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := SumIntegersInString(test.input)
			if got != test.want {
				t.Errorf("For string %q: got %d, want %d", test.input, got, test.want)
			}
		})
	}
}
