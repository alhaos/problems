package tickets

import (
	"testing"
)

func TestAssemble(t *testing.T) {

	testCases := []struct {
		name     string
		count    int
		expected ticketSet
	}{
		{
			name:  "main test",
			count: 131,
			expected: ticketSet{
				x1:  1,
				x5:  0,
				x10: 1,
				x20: 0,
				x60: 2,
			},
		},
		{
			name:  "test #2",
			count: 120,
			expected: ticketSet{
				x1:  0,
				x5:  0,
				x10: 0,
				x20: 0,
				x60: 2,
			},
		},
	}

	for _, tC := range testCases {

		t.Run(tC.name, func(t *testing.T) {
			result := assemble(tC.count)
			if result != tC.expected {
				t.Errorf("unexpected result for test %s, expected: %v, but got: %v", tC.name, tC.expected, result)
			}
		})
	}

}
