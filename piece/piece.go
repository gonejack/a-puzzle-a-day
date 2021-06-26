package piece

import (
	"fmt"
	"strings"

	"github.com/gonejack/a-puzzle-a-day/board"
)

var pieces = []piece{
	{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
	},
	{
		{1, 0, 0},
		{1, 1, 1},
		{0, 0, 1},
	},
	{
		{0, 0, 1, 1},
		{1, 1, 1, 0},
	},
	{
		{1, 1, 1, 1},
		{0, 0, 1, 0},
	},
	{
		{0, 1, 1},
		{1, 1, 1},
	},
	{
		{1, 1, 1},
		{1, 0, 1},
	},
	{
		{1, 1, 1},
		{1, 1, 1},
	},
	{
		{1, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	},
}

var Pieces [][]piece

func init() {
	for _, p := range pieces {
		pt := p.transforms()
		Pieces = append(Pieces, pt)
	}
}

const n = 4

type piece [n][n]int

func (p piece) Print() {
	for x := range p {
		fmt.Println(p[x])
	}
	fmt.Println(strings.Repeat("-", 10))
}
func (p piece) rotate() (r piece) {
	for x := range r {
		for y := range r[x] {
			r[x][y] = p[y][n-x-1]
		}
	}
	return
}
func (p piece) flip() (f piece) {
	for x := range f {
		for i := 0; i < n; i++ {
			f[x][i] = p[x][n-i-1]
		}
	}
	return
}
func (p piece) transforms() (ps []piece) {
	ps = append(ps, p)

	for i := 0; i < 3; i++ {
		ps = append(ps, ps[i].rotate())
	}
	for i := 4; i < 8; i++ {
		ps = append(ps, ps[i-4].flip())
	}

	m := make(map[piece]struct{})
	for _, p := range ps {
		m[p.shift()] = struct{}{}
	}

	ps = nil
	for p := range m {
		ps = append(ps, p)
	}

	return
}
func (p piece) shift() (s piece) {
	si := 0
	for _, r := range p {
		if r[0]+r[1]+r[2]+r[3] == 0 {
			continue
		}
		s[si] = r
		si += 1
	}

	for {
		for i := range s {
			if s[i][0] != 0 {
				return
			}
		}

		for i := range s {
			for j := range s[i] {
				if j < 3 {
					s[i][j] = s[i][j+1]
				} else {
					s[i][j] = 0
				}
			}
		}
	}
}
func (p piece) CanPlace(b *board.Board7x7, x, y int) bool {
	return p.put(b, x, y, "", false)
}
func (p piece) Place(b *board.Board7x7, x, y int, text string) bool {
	return p.put(b, x, y, text, true)
}
func (p piece) put(b *board.Board7x7, x int, y int, text string, doWrite bool) (suc bool) {
	for i := range p {
		for j := range p[i] {
			if p[i][j] == 0 {
				continue
			}

			tx := x + i
			ty := y + j
			ok := b.CanSet(tx, ty)
			if !ok {
				return false
			}
			if doWrite {
				b.Set(text, tx, ty)
			}
		}
	}

	return true
}
