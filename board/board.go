package board

import (
	"fmt"
	"strings"
)

type Board7x7 [7][7]struct {
	text string
	used bool
}

func (b *Board7x7) reset() *Board7x7 {
	const dat7x7 = `
	JAN	|	FEB	|	MAR	|	APR	|	MAY	|	JUN	|	-
	JUL	|	AUG	|	SEP	|	OCT	|	NOV	|	DEC	|	-
	1	|	2	|	3	|	4	|	5	|	6	|	7
	8	|	9	|	10	|	11	|	12	|	13 	|	14
	15	|	16	|	17	|	18	|	19	|	20 	|	21
	22	|	23	|	24	|	25	|	26	|	27 	|	28
	29	|	30	|	31	|	-	|	-	|	- 	|	-
`
	for r, line := range strings.Split(strings.Trim(dat7x7, "\n"), "\n") {
		for c, word := range strings.Split(line, "|") {
			w := strings.TrimSpace(word)
			b[r][c].text = w
			b[r][c].used = w == "-"
		}
	}
	return b
}
func (b *Board7x7) setDate(month string, day string) *Board7x7 {
	for i := range b {
		for j := range b[i] {
			switch {
			case strings.EqualFold(b[i][j].text, month):
				b[i][j].used = true
			case strings.EqualFold(b[i][j].text, day):
				b[i][j].used = true
			}
		}
	}
	return b
}

func (b *Board7x7) Clone() *Board7x7 {
	c := *b
	return &c
}
func (b *Board7x7) CanSet(r int, c int) (ok bool) {
	if r < 0 || r > 6 || c < 0 || c > 6 {
		return false
	}
	return b[r][c].used == false
}
func (b *Board7x7) Set(Text string, r int, c int) {
	b[r][c].text = Text
	b[r][c].used = true
}
func (b *Board7x7) Print() {
	for r := range b {
		for c := range b[r] {
			text := b[r][c].text
			tlen := len(text)
			fmt.Print(text)
			switch {
			case tlen == 3:
				fmt.Print(" ")
			case tlen == 2:
				fmt.Print(" ")
			case c+1 < len(b[r]) && len(b[r][c+1].text) == 3:
				fmt.Print(" ")
			default:
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", 20))
}

func NewBoard7x7(mon string, day string) *Board7x7 {
	var b Board7x7
	return b.reset().setDate(mon, day)
}
