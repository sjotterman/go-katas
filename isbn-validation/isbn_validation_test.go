package kata

import "testing"

func TestValidateISBN(t *testing.T) {

	type test struct {
		input string
		want  bool
	}

	tests := []test{
		{"1112223339", true},
		{"048665088X", true},
		{"1293000000", true},
		{"1234554321", true},
		{"1234512345", false},
		{"048665088X", true},
		{"X123456788", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := ValidateISBN(test.input)
			if got != test.want {
				t.Errorf("For value %q: got %v, want %v", test.input, got, test.want)
			}
		})
	}
}
