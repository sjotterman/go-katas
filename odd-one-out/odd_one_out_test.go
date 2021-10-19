package kata

import (
	"reflect"
	"testing"
)

func TestOddOneOut(t *testing.T) {

	type test struct {
		name  string
		input []int
		want  []int
	}

	tests := []test{
		{name: "3 is out", input: []int{2, 3, 4, 6, 8, 10}, want: []int{2, 4, 6, 8, 10}},
		{name: "1 is out", input: []int{6, 1, 4, 22, 104}, want: []int{6, 4, 22, 104}},
		{name: "several even are out", input: []int{1, 2, 3, 4, 6, 7, 9, 11, 13, 15}, want: []int{1, 3, 7, 9, 11, 13, 15}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := OddOneOut(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("For value %v: got %v, want %v", test.input, got, test.want)
			}
		})
	}
}
