package wordsearch

import (
	"slices"
	"testing"
)

func makeBoard(rows ...string) [][]byte {
	board := make([][]byte, len(rows))
	for i, r := range rows {
		board[i] = []byte(r)
	}
	return board
}

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{"example 1", makeBoard("ABCE", "SFCS", "ADEE"), "ABCCED", true},
		{"example 2", makeBoard("ABCE", "SFCS", "ADEE"), "SEE", true},
		{"example 3", makeBoard("ABCE", "SFCS", "ADEE"), "ABCB", false},
		{"single cell match", makeBoard("a"), "a", true},
		{"single cell miss", makeBoard("a"), "b", false},
		{"cell reuse forbidden", makeBoard("aa"), "aaa", false},
		{"around the square", makeBoard("ab", "cd"), "abdc", true},
		{"no diagonal moves", makeBoard("ab", "cd"), "abcd", false},
		{"backtrack needed", makeBoard("CAA", "AAA", "BCD"), "AAB", true},
		{"word longer than board", makeBoard("ab", "cd"), "abcdabcd", false},
		{"single row", makeBoard("abcde"), "edcba", true},
		{"single column", makeBoard("a", "b", "c"), "cb", true},
		{"missing letter", makeBoard("abc", "def"), "abz", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := make([][]byte, len(tt.board))
			for i := range tt.board {
				orig[i] = slices.Clone(tt.board[i])
			}

			got := exist(tt.board, tt.word)

			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}

			for i := range orig {
				if !slices.Equal(orig[i], tt.board[i]) {
					t.Fatalf("board mutated: row %d is %q, was %q", i, tt.board[i], orig[i])
				}
			}
		})
	}
}
