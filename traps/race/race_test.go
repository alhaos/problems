package race

import "testing"

func TestCounterWithRace(t *testing.T) {
	testCases := []struct {
		desc     string
		n        int
		expected int
	}{
		{
			desc:     "1 итерация",
			n:        1,
			expected: 1,
		},
		{
			desc:     "10 итераций",
			n:        10,
			expected: 10,
		},
		{
			desc:     "100 итераций",
			n:        100,
			expected: 100,
		},
		{
			desc:     "1000 итераций",
			n:        1000,
			expected: 1000,
		},
		{
			desc:     "10000 итераций",
			n:        10000,
			expected: 10000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := CounterWithRace(tC.n)
			if result != tC.expected {
				t.Errorf("unexpected result for test %s, expected: %d, but got: %d", tC.desc, tC.expected, result)
			}
		})
	}
}

func TestCounterWithAtomic(t *testing.T) {
	testCases := []struct {
		desc     string
		n        int
		expected int
	}{
		{
			desc:     "1 итерация",
			n:        1,
			expected: 1,
		},
		{
			desc:     "10 итераций",
			n:        10,
			expected: 10,
		},
		{
			desc:     "100 итераций",
			n:        100,
			expected: 100,
		},
		{
			desc:     "1000 итераций",
			n:        1000,
			expected: 1000,
		},
		{
			desc:     "10000 итераций",
			n:        10000,
			expected: 10000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := CounterWithAtomic(tC.n)
			if result != tC.expected {
				t.Errorf("unexpected result for test %s, expected: %d, but got: %d", tC.desc, tC.expected, result)
			}
		})
	}
}

func TestCounterWithMutex(t *testing.T) {
	testCases := []struct {
		desc     string
		n        int
		expected int
	}{
		{
			desc:     "1 итерация",
			n:        1,
			expected: 1,
		},
		{
			desc:     "10 итераций",
			n:        10,
			expected: 10,
		},
		{
			desc:     "100 итераций",
			n:        100,
			expected: 100,
		},
		{
			desc:     "1000 итераций",
			n:        1000,
			expected: 1000,
		},
		{
			desc:     "10000 итераций",
			n:        10000,
			expected: 10000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := CounterWithMutex(tC.n)
			if result != tC.expected {
				t.Errorf("unexpected result for test %s, expected: %d, but got: %d", tC.desc, tC.expected, result)
			}
		})
	}
}
