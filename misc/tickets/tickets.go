package tickets

// ticket catalog: trips covered and price
type ticketType struct {
	trips int
	price int
	name  string
}

var catalog = []ticketType{
	{1, 30, "x1"},
	{5, 120, "x5"},
	{10, 235, "x10"},
	{20, 420, "x20"},
	{60, 1005, "x60"},
}

// small ticket types used to cover the "remainder" after fixing how many
// x60 tickets we buy. 60 is deliberately excluded here: the number of x60
// tickets is chosen explicitly in solve() via a small search window.
var smallCatalog = []ticketType{
	{1, 30, "x1"},
	{5, 120, "x5"},
	{10, 235, "x10"},
	{20, 420, "x20"},
}

// remainderTable is precomputed once: for r trips (0..maxRemainder),
// minCost[r] = cheapest way to cover >= r trips using only x1/x5/x10/x20,
// choice[r] = which ticket type was used last (for reconstruction).
const maxRemainder = 130

var minCost [maxRemainder + 1]int
var choice [maxRemainder + 1]int // index into smallCatalog

func init() {
	const inf = 1 << 30
	for r := 1; r <= maxRemainder; r++ {
		best := inf
		bestIdx := -1
		for i, t := range smallCatalog {
			prev := r - t.trips
			if prev < 0 {
				prev = 0
			}
			c := t.price + minCost[prev]
			if c < best {
				best = c
				bestIdx = i
			}
		}
		minCost[r] = best
		choice[r] = bestIdx
	}
}

// reconstructSmall returns ticket counts for x1,x5,x10,x20 covering >= r trips
// at minimum cost, using the precomputed table above.
func reconstructSmall(r int) (x1, x5, x10, x20 int) {
	for r > 0 {
		i := choice[r]
		t := smallCatalog[i]
		switch t.trips {
		case 1:
			x1++
		case 5:
			x5++
		case 10:
			x10++
		case 20:
			x20++
		}
		r -= t.trips
		if r < 0 {
			r = 0
		}
	}
	return
}

// solve finds the minimum-cost ticket combination covering at least n trips.
// n can be up to 1e9.
func solve(n int) ticketSet {
	if n <= 0 {
		return ticketSet{}
	}
	const x60price = 1005
	jBase := n / 60
	jLo := jBase - 2
	if jLo < 0 {
		jLo = 0
	}
	jHi := jBase + 2

	bestCost := -1
	var best ticketSet

	for j := jLo; j <= jHi; j++ {
		rem := n - 60*j
		if rem < 0 {
			rem = 0
		}
		if rem > maxRemainder {
			continue // shouldn't happen given the window size
		}
		total := j*x60price + minCost[rem]
		if bestCost == -1 || total < bestCost {
			x1, x5, x10, x20 := reconstructSmall(rem)
			bestCost = total
			best = ticketSet{x1: x1, x5: x5, x10: x10, x20: x20, x60: j}
		}
	}
	return best
}
