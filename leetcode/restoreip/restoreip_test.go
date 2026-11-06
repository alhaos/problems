package restoreip

import (
	"slices"
	"testing"
)

func TestRestoreIPAddresses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"two results", "25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"all zeros", "0000", []string{"0.0.0.0"}},
		{"five results", "101023", []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
		{"all ones", "1111", []string{"1.1.1.1"}},
		{"leading zeros", "010010", []string{"0.10.0.10", "0.100.1.0"}},
		{"max octets", "255255255255", []string{"255.255.255.255"}},
		{"too short", "123", nil},
		{"too long", "1234567890123", nil},
		{"octet over 255", "256256256256", nil},
		{"only invalid splits", "000256", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RestoreIPAddresses(tt.s)
			slices.Sort(got)
			slices.Sort(tt.want)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("RestoreIPAddresses(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
