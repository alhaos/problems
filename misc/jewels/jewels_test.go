package jewels

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		name   string
		jewels string
		stones string
		want   int
	}{
		{
			name:   "базовый пример",
			jewels: "aA",
			stones: "aAAbbbb",
			want:   3,
		},
		{
			name:   "нет драгоценностей среди камней",
			jewels: "z",
			stones: "aAAbbbb",
			want:   0,
		},
		{
			name:   "все камни — драгоценности",
			jewels: "abc",
			stones: "abcabc",
			want:   6,
		},
		{
			name:   "регистр имеет значение",
			jewels: "a",
			stones: "aAAAA",
			want:   1,
		},
		{
			name:   "один камень, одна драгоценность — совпадение",
			jewels: "x",
			stones: "x",
			want:   1,
		},
		{
			name:   "один камень, одна драгоценность — не совпадение",
			jewels: "x",
			stones: "y",
			want:   0,
		},
		{
			name:   "много драгоценностей, один тип камней",
			jewels: "abcdefg",
			stones: "aaaaaaa",
			want:   7,
		},
		{
			name:   "только заглавные буквы",
			jewels: "AB",
			stones: "AABBCC",
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numJewelsInStones(tt.jewels, tt.stones)
			if got != tt.want {
				t.Errorf("numJewelsInStones(%q, %q) = %d, want %d",
					tt.jewels, tt.stones, got, tt.want)
			}
		})
	}
}
