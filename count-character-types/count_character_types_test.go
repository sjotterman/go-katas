package kata

import (
	"reflect"
	"testing"
)

func TestCountCharacters(t *testing.T) {

	type test struct {
		name  string
		input string
		want  []int
	}

	tests := []test{
		{"first", "Codewars@codewars123.com", []int{1, 18, 3, 2}},
		{"second", "bgA5<1d-tOwUZTS8yQ", []int{7, 6, 3, 2}},
		{"third", "P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H", []int{9, 9, 6, 9}},
		{"fourth", "RYT'>s&gO-.CM9AKeH?,5317tWGpS<*x2ukXZD", []int{15, 8, 6, 9}},
		{"fifth", "$Cnl)Sr<7bBW-&qLHI!mY41ODe", []int{10, 7, 3, 6}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CountCharacters(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("For value %v: got %v, want %v", test.input, got, test.want)
			}
		})
	}
}
