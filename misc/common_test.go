package misc

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		desc     string
		a        []int
		b        []int
		expected []int
	}{
		{
			desc:     "#1",
			a:        []int{23, 3, 1, 2},
			b:        []int{6, 2, 4, 23},
			expected: []int{2, 23},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := intersection(tC.a, tC.b)
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("Unexpected result for test case [%s] for values: [%+v, %+v] expected: [%+v], but got [%+v] ", tC.desc, tC.a, tC.b, tC.expected, result)
			}
			t.Logf("Expected result for test case [%s] for values: [%+v, %+v] expected: [%+v], and got [%+v] ", tC.desc, tC.a, tC.b, tC.expected, result)
		})
	}
}
