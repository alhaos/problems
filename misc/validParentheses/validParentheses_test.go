package validParentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "одна пара круглых скобок",
			s:    "()",
			want: true,
		},
		{
			name: "несколько пар разных скобок подряд",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "вложенные скобки",
			s:    "{[]}",
			want: true,
		},
		{
			name: "неправильный тип закрывающей скобки",
			s:    "(]",
			want: false,
		},
		{
			name: "правильно открыты, неправильно закрыты",
			s:    "([)]",
			want: false,
		},
		{
			name: "только открывающие скобки",
			s:    "(((",
			want: false,
		},
		{
			name: "только закрывающие скобки",
			s:    ")))",
			want: false,
		},
		{
			name: "нечётная длина строки",
			s:    "([]",
			want: false,
		},
		{
			name: "глубокая вложенность",
			s:    "{[()]}",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
