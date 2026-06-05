package brutForce

type Alphabet []byte

type BrutForce interface {
	Generate(alphabet []byte, from int, to int) []string
}

type brutForce struct {
	cases []string
}

func NewBrutForce(alphabet string) BrutForce {
	return &brutForce{}
}

func (bf *brutForce) Generate(alphabet []byte, from int, to int) []string {
	l := len(alphabet)
	for n := from; n <= to; n++ {
		indexes := make([]int, n)

		for i := n - 1; i >= 0; i-- {
			if indexes[i] < l {
				indexes[i]++
			} else {
				indexes[i] = 0
				indexes[i-1]++
			}
		}
	}

	return []string{""}
}
