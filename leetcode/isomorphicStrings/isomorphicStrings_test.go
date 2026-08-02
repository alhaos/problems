package isomorphic

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		target string
		want   bool
	}{
		// {
		// 	name:   "egg_add",
		// 	s:      "egg",
		// 	target: "add",
		// 	want:   true,
		// },
		// {
		// 	name:   "foo_bar",
		// 	s:      "foo",
		// 	target: "bar",
		// 	want:   false,
		// },
		// {
		// 	name:   "paper_title",
		// 	s:      "paper",
		// 	target: "title",
		// 	want:   true,
		// },
		// {
		// 	name:   "badc_baba",
		// 	s:      "badc",
		// 	target: "baba",
		// 	want:   false,
		// },
		// {
		// 	name:   "ab_aa",
		// 	s:      "ab",
		// 	target: "aa",
		// 	want:   false,
		// },
		// {
		// 	name:   "empty_strings",
		// 	s:      "",
		// 	target: "",
		// 	want:   true,
		// },
		// {
		// 	name:   "single_char",
		// 	s:      "a",
		// 	target: "a",
		// 	want:   true,
		// },
		{
			name:   "ab_ca",
			s:      "ab",
			target: "ca",
			want:   true,
		},
		{
			name:   "aa_ab",
			s:      "aa",
			target: "ab",
			want:   false,
		},
		{
			name:   "different_length",
			s:      "abcdefghijklmnop",
			target: "truevaluesarehere",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isIsomorphic(tt.s, tt.target)
			if got != tt.want {
				t.Errorf("isIsomorphic(%q, %q) = %v, want %v", tt.s, tt.target, got, tt.want)
			}
		})
	}
}
