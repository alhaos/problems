package jewels

func numJewelsInStones(jewels string, stones string) int {

	jewelsMap := make(map[rune]struct{})

	for _, jewel := range jewels {
		jewelsMap[jewel] = struct{}{}
	}

	counter := 0

	for _, stone := range stones {
		if _, exist := jewelsMap[stone]; exist {
			counter++
		}
	}

	return counter

}
