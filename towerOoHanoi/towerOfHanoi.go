package towerOoHanoi

import "fmt"

type Disk int

type PegIndex int

const (
	pegA PegIndex = iota
	pegB
	pegC
)

type Pegs [3]int

func (p *Pegs) Init(n int) {
	p[pegA] = n
}

func (p *Pegs) Solve(n int, from, to, aux PegIndex) {
	if n == 1 {
		fmt.Printf("Переместить диск со стержня %c на стержень %c\n", 'A'+byte(from), 'A'+byte(to))
		return
	}

	p.Solve(n-1, from, aux, to)

	p.Solve(1, from, to, aux)

	p.Solve(n-1, aux, to, from)
}

func (p *Pegs) SolveTower(n int) {
	p.Solve(n, pegA, pegC, pegB)
}

func (p *Pegs) Print() {
	fmt.Printf("Стержень A: %d дисков, B: %d дисков, C: %d дисков\n", p[pegA], p[pegB], p[pegC])
}

func (p *Pegs) IsSolved(n int) bool {
	return p[pegC] == n
}
