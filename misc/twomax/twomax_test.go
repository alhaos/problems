package twomax

import "testing"

func TestTwoMax(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		wantMax1 int
		wantMax2 int
		wantOk   error
	}{
		{"basic", []int{3, 1, 4, 1, 5, 9, 2, 6}, 9, 6, nil},
		{"sorted asc", []int{1, 2, 3, 4, 5}, 5, 4, nil},
		{"sorted desc", []int{5, 4, 3, 2, 1}, 5, 4, nil},
		{"duplicates at top", []int{7, 7, 3, 1}, 7, 7, nil},
		{"negative numbers", []int{-5, -1, -3, -2}, -1, -2, nil},
		{"two elements", []int{10, 20}, 20, 10, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			max1, max2, err := TwoMax(c.nums)
			if err != nil {
				t.Error(err)
			}
			if max1 != c.wantMax1 || max2 != c.wantMax2 {
				t.Errorf("got (%d, %d), want (%d, %d)", max1, max2, c.wantMax1, c.wantMax2)
			}
		})
	}
}
