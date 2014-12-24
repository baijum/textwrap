package textwrap

import (
	"reflect"
	"testing"
)

func TestWrap(t *testing.T) {
	var tests = []struct {
		text  string
		width int
		out   []string
	}{
		{"hello", 2, []string{"he", "ll", "o"}},
		{"good", 2, []string{"go", "od"}},
		{"hello", 3, []string{"hel", "lo"}},
	}
	for _, test := range tests {
		actual := Wrap(test.text, test.width)
		if !reflect.DeepEqual(actual, test.out) {
			t.Errorf("Wrap(%q, %d) = %#v; want %#v",
				test.text, test.width, actual, test.out)
		}
	}
}
