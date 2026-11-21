package medianfinder

import "testing"

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name string
		adds []int
		want []float64
	}{
		{"example", []int{1, 2, 3}, []float64{1, 1.5, 2}},
		{"single", []int{5}, []float64{5}},
		{"descending", []int{5, 4, 3, 2, 1}, []float64{5, 4.5, 4, 3.5, 3}},
		{"ascending", []int{1, 2, 3, 4, 5, 6}, []float64{1, 1.5, 2, 2.5, 3, 3.5}},
		{"all same", []int{2, 2, 2, 2}, []float64{2, 2, 2, 2}},
		{"negatives", []int{-1, -2, -3}, []float64{-1, -1.5, -2}},
		{"mixed", []int{41, 35, 62, 4, 97, 108}, []float64{41, 38, 41, 38, 41, 51.5}},
		{"with zeros", []int{0, 0, 1, 2}, []float64{0, 0, 0, 0.5}},
		{"big jumps", []int{100000, -100000, 0}, []float64{100000, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := NewMedianFinder()
			for i, v := range tt.adds {
				mf.AddNum(v)
				if got := mf.FindMedian(); got != tt.want[i] {
					t.Errorf("after AddNum(%d) #%d: got %v, want %v", v, i, got, tt.want[i])
				}
			}
		})
	}
}
