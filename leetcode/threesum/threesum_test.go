package threesum

import (
	"math/rand"
	"slices"
	"testing"
)

// normalize сортирует числа внутри троек и сами тройки,
// чтобы сравнение не зависело от порядка.
func normalize(triples [][]int) [][]int {
	res := make([][]int, 0, len(triples))
	for _, tr := range triples {
		c := slices.Clone(tr)
		slices.Sort(c)
		res = append(res, c)
	}
	slices.SortFunc(res, func(a, b []int) int { return slices.Compare(a, b) })
	return res
}

func equalTriples(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func hasDuplicates(sorted [][]int) bool {
	for i := 1; i < len(sorted); i++ {
		if slices.Equal(sorted[i], sorted[i-1]) {
			return true
		}
	}
	return false
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"example1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"no answer", []int{0, 1, 1}, [][]int{}},
		{"all zeros", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"many zeros", []int{0, 0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"all positive", []int{1, 2, 3, 4}, [][]int{}},
		{"all negative", []int{-1, -2, -3, -4}, [][]int{}},
		{"exactly three", []int{-1, 0, 1}, [][]int{{-1, 0, 1}}},
		{"duplicates collapse", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"same value twice in triple", []int{-4, 2, 2, 2}, [][]int{{-4, 2, 2}}},
		{"two triples with shared number", []int{-2, -2, 0, 2, 2, 4, -4}, [][]int{{-4, 0, 4}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
		{"sorted input", []int{-3, -1, 0, 1, 2, 4}, [][]int{{-3, -1, 4}, {-3, 1, 2}, {-1, 0, 1}}},
		{"unsorted input", []int{4, -1, 2, 1, -3, 0}, [][]int{{-3, -1, 4}, {-3, 1, 2}, {-1, 0, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(threeSum(slices.Clone(tt.nums)))
			want := normalize(tt.want)

			if hasDuplicates(got) {
				t.Errorf("threeSum(%v) has duplicate triples: %v", tt.nums, got)
			}
			if !equalTriples(got, want) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, want)
			}
		})
	}
}

// brute перебирает все тройки индексов и собирает уникальные по значениям.
func brute(nums []int) [][]int {
	seen := make(map[[3]int]struct{})
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tr := []int{nums[i], nums[j], nums[k]}
					slices.Sort(tr)
					seen[[3]int{tr[0], tr[1], tr[2]}] = struct{}{}
				}
			}
		}
	}
	res := make([][]int, 0, len(seen))
	for k := range seen {
		res = append(res, []int{k[0], k[1], k[2]})
	}
	return res
}

func TestThreeSumRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < 20000; i++ {
		nums := make([]int, rng.Intn(10)+3)
		for j := range nums {
			nums[j] = rng.Intn(9) - 4
		}

		got := normalize(threeSum(slices.Clone(nums)))
		want := normalize(brute(nums))

		if hasDuplicates(got) {
			t.Fatalf("nums=%v: duplicate triples in %v", nums, got)
		}
		if !equalTriples(got, want) {
			t.Fatalf("nums=%v got=%v want=%v", nums, got, want)
		}
	}
}
