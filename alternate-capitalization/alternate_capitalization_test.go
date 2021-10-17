package kata

import (
	"reflect"
	"testing"
)

func TestAlternateCapitalization(t *testing.T) {

	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{input: "ab", want: []string{"Ab", "aB"}},
		{input: "codewars", want: []string{"CoDeWaRs", "cOdEwArS"}},
		{input: "abracadabra", want: []string{"AbRaCaDaBrA", "aBrAcAdAbRa"}},
		{input: "codewarriors", want: []string{"CoDeWaRrIoRs", "cOdEwArRiOrS"}},
		{input: "indexinglessons", want: []string{"InDeXiNgLeSsOnS", "iNdExInGlEsSoNs"}},
		{input: "codingisafunactivity", want: []string{"CoDiNgIsAfUnAcTiViTy", "cOdInGiSaFuNaCtIvItY"}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := ToAlternateCapitalization(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("For value %q: got %v, want %v", test.input, got, test.want)
			}
		})
	}
}
