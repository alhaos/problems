package codewars

import "testing"

func TestRoundToNext5(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected int
	}{
		{"with_0", 0, 0},
		{"with_2", 2, 5},
		{"with_3", 3, 5},
		{"with_12", 12, 15},
		{"with_21", 21, 25},
		{"with_30", 30, 30},
		{"with_-2", -2, 0},
		{"with_-5", -5, -5},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := RoundToNext5(tC.n)
			if result != tC.expected {
				t.Errorf("Unexpected result from test %s when n=[%d] expected: [%d], but got [%d]", tC.name, tC.n, tC.expected, result)
			}
		})
	}
}

func TestSummation(t *testing.T) {
	testCases := []struct {
		desc     string
		n        int
		expected int
	}{
		{
			desc:     "with 2",
			n:        2,
			expected: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := Summation(tC.n)
			if result != tC.expected {
				t.Errorf("Unexpected result from test %s when n=[%d] expected: [%d], but got [%d]", tC.desc, tC.n, tC.expected, result)
			}
		})
	}
}

func TestEncryptThis(t *testing.T) {
	testCases := []struct {
		desc     string
		s        string
		expected string
	}{
		{
			desc:     "with_Hello",
			s:        "Hello",
			expected: "72olle",
		},
		{
			desc:     "with_good",
			s:        "good",
			expected: "103doo",
		},
		{
			desc:     "with_hello_world",
			s:        "hello world",
			expected: "104olle 119drlo",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := EncryptThis(tC.s)
			if result != tC.expected {
				t.Errorf("Unexpected result from test %s when s=[%s] expected: [%s], but got [%s]", tC.desc, tC.s, tC.expected, result)
			}
		})
	}
}
