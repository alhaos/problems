package meetingrooms

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		want int
	}{
		{"basic", [][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{"unsorted no overlap", [][]int{{7, 10}, {2, 4}}, 1},
		{"touching", [][]int{{1, 5}, {5, 10}}, 1},
		{"single", [][]int{{1, 2}}, 1},
		{"all same", [][]int{{1, 5}, {1, 5}, {1, 5}}, 3},
		{"disjoint", [][]int{{1, 2}, {3, 4}, {5, 6}}, 1},
		{"nested small ones sequential", [][]int{{1, 10}, {2, 3}, {4, 5}}, 2},
		{"staircase", [][]int{{1, 4}, {2, 5}, {3, 6}, {4, 7}}, 3},
		{"unsorted overlap", [][]int{{9, 10}, {4, 9}, {4, 17}}, 2},
		{"one long, many short", [][]int{{0, 5}, {1, 2}, {3, 4}}, 2},
		{"mixed", [][]int{{2, 15}, {36, 45}, {9, 29}, {16, 23}, {4, 9}}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinMeetingRooms(tt.in)
			if got != tt.want {
				t.Errorf("MinMeetingRooms(%v) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
