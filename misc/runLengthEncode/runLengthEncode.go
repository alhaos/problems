package runLengthEncode

import (
	"fmt"
	"strings"
	"unicode"
)

type prevChar struct {
	char  rune
	count int
}

func (pc *prevChar) String() string {
	if pc.count == 1 {
		return string(pc.char)
	}
	return fmt.Sprintf("%d%c", pc.count, pc.char)
}

// Encode compresses a string using run-length encoding.
// Single characters are not prefixed with "1".
// Example: "AAABCCDDDD" -> "3AB2C4D"
func Encode(s string) string {
	switch len(s) {
	case 0, 1:
		return s
	}
	var sb strings.Builder
	var pc prevChar
	for i, char := range s {
		switch i {
		case 0:
			pc.char = char
			pc.count = 1
		default:
			if char == pc.char {
				pc.count++
			} else {
				sb.WriteString(pc.String())
				pc.char = char // ← запоминаем новый символ
				pc.count = 1
			}
		}
	}
	sb.WriteString(pc.String())
	return sb.String()
}

// Decode expands a run-length encoded string.
// Example: "3AB2C4D" -> "AAABCCDDDD"
func Decode(s string) string {
	factor := 0
	var sb strings.Builder
	for _, char := range s {
		if unicode.IsDigit(char) {
			factor = factor*10 + int(char-'0')
		} else {
			if factor == 0 {
				factor = 1
			}
			sb.WriteString(strings.Repeat(string(char), factor))
			factor = 0
		}
	}
	return sb.String()
}
