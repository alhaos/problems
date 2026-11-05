package acmp39

//
func HairyBusiness(prices []int) int {

	var s, m int

	for i := len(prices) - 1; i >= 0; i-- {
		m = max(m, prices[i])
		s += m
	}

	return s
}
