package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "			",
			expected: []string{},
		},
		{
			input:    "		hello	",
			expected: []string{"hello"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			if word != c.expected[i] {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}

func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestParseInput(t *testing.T) {
	cases := []struct {
		input    string
		wantCmd  string
		wantArgs []string
	}{
		{
			input:    "map",
			wantCmd:  "map",
			wantArgs: []string{}, // or nil, depending on your implementation
		},
		{
			input:    "explore somewhere",
			wantCmd:  "explore",
			wantArgs: []string{"somewhere"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			cmdString, args := parseInput(tc.input)
			if tc.wantCmd != cmdString {
				t.Errorf("For input '%s', expected command '%s', but got '%s'", tc.input, tc.wantCmd, cmdString)
			}
			if !compareStringSlices(args, tc.wantArgs) {
				t.Errorf("For input '%s', expected command '%s', but got '%s'", tc.input, tc.wantCmd, cmdString)
			}
		})
	}
}
