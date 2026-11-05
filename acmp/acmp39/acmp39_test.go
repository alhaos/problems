package acmp39

import "testing"

func TestHairyBusiness(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		// примеры с acmp.ru
		{"example 1", []int{73, 31, 96, 24, 46}, 380},
		{"example 2: increasing", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100},
		{"example 3: decreasing", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 55},

		// граничные случаи
		{"single day", []int{5}, 5},
		{"single day max price", []int{100}, 100},
		{"two days, second better", []int{1, 2}, 4},
		{"two days, first better", []int{2, 1}, 3},
		{"all equal", []int{5, 5, 5}, 15},

		// смешанные случаи
		{"peak in the middle", []int{1, 100, 1}, 201},
		{"peak at start", []int{100, 1, 1}, 102},
		{"sell first, wait for last", []int{3, 1, 2}, 7},
		{"spike in the middle", []int{1, 1, 50, 1, 1}, 152},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HairyBusiness(tt.prices)
			if got != tt.want {
				t.Errorf("HairyBusiness(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}

func TestHairyBusinessMaxN(t *testing.T) {
	const n = 100

	t.Run("all 100", func(t *testing.T) {
		prices := make([]int, n)
		for i := range prices {
			prices[i] = 100
		}
		if got, want := HairyBusiness(prices), 10000; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("increasing 1..100", func(t *testing.T) {
		prices := make([]int, n)
		for i := range prices {
			prices[i] = i + 1
		}
		if got, want := HairyBusiness(prices), 10000; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("decreasing 100..1", func(t *testing.T) {
		prices := make([]int, n)
		for i := range prices {
			prices[i] = n - i
		}
		if got, want := HairyBusiness(prices), 5050; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
