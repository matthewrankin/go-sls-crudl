package parse

import (
	"testing"
)

func TestUnslugify(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"foo", "foo"},
		{"foo-bar", "foo bar"},
		{"foo-bar+baz", "foo bar baz"},
	}
	for _, test := range tests {
		if got := Unslugify(test.input); got != test.want {
			t.Errorf("Unslugify(%q) = %v", test.input, got)
		}
	}
}

func TestSlugify(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"foo", "foo"},
		{"foo bar", "foo-bar"},
		{"foo bar baz", "foo-bar-baz"},
	}
	for _, test := range tests {
		if got := Slugify(test.input); got != test.want {
			t.Errorf("Slugify(%q) = %v", test.input, got)
		}
	}
}
