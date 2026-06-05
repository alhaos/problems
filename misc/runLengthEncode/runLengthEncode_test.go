package runLengthEncode

import "testing"

var encodeTests = []struct {
	name  string
	input string
	want  string
}{
	{"empty string", "", ""},
	{"single char", "A", "A"},
	{"no repeats", "ABC", "ABC"},
	{"all same", "AAAA", "4A"},
	{"basic example", "AAABCCDDDD", "3AB2C4D"},
	{"mixed case", "aaBBcc", "2a2B2c"},
	{"single then run", "XYYY", "X3Y"},
	{"spaces", "  A  B", "2 A2 B"},
	{"long run", "ZZZZZZZZZZ", "10Z"},
	{"unicode letters", "ааббб", "2а3б"},
}

var decodeTests = []struct {
	name  string
	input string
	want  string
}{
	{"empty string", "", ""},
	{"no numbers", "ABC", "ABC"},
	{"basic decode", "3AB2C4D", "AAABCCDDDD"},
	{"single after run", "2XY", "XXY"},
	{"two-digit count", "10Z", "ZZZZZZZZZZ"},
	{"spaces encoded", "2 A2 B", "  A  B"},
	{"roundtrip check", "X3Y", "XYYY"},
}

func TestEncode(t *testing.T) {
	for _, tt := range encodeTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.input)
			if got != tt.want {
				t.Errorf("Encode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for _, tt := range decodeTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Decode(tt.input)
			if got != tt.want {
				t.Errorf("Decode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestRoundtrip(t *testing.T) {
	inputs := []string{
		"AAABCCDDDD",
		"Hello, World!",
		"aabbccdd",
		"ABCDEFG",
		"     ",
	}
	for _, s := range inputs {
		got := Decode(Encode(s))
		if got != s {
			t.Errorf("Decode(Encode(%q)) = %q, want original", s, got)
		}
	}
}
