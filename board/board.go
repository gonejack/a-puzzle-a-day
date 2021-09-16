package board

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var Board = Board7x7{
	{{"JAN", 0}, {"FEB", 0}, {"MAR", 0}, {"APR", 0}, {"MAY", 0}, {"JUN", 0}, {"", 1}},
	{{"JUL", 0}, {"AUG", 0}, {"SEP", 0}, {"OCT", 0}, {"NOV", 0}, {"DEC", 0}, {"", 1}},
	{{"1", 0}, {"2", 0}, {"3", 0}, {"4", 0}, {"5", 0}, {"6", 0}, {"7", 0}},
	{{"8", 0}, {"9", 0}, {"10", 0}, {"11", 0}, {"12", 0}, {"13", 0}, {"14", 0}},
	{{"15", 0}, {"16", 0}, {"17", 0}, {"18", 0}, {"19", 0}, {"20", 0}, {"21", 0}},
	{{"22", 0}, {"23", 0}, {"24", 0}, {"25", 0}, {"26", 0}, {"27", 0}, {"28", 0}},
	{{"29", 0}, {"30", 0}, {"31", 0}, {"", 1}, {"", 1}, {"", 1}, {"", 1}},
}

type (
	cell struct {
		Text string
		Flag int
	}
	Board7x7 [7][7]cell
)

func (b *Board7x7) CanSet(r int, c int) (ok bool) {
	if r < 0 || r > 6 || c < 0 || c > 6 {
		return false
	}
	return b[r][c].Flag == 0
}
func (b *Board7x7) Set(Text string, r int, c int) {
	b[r][c].Text = Text
	b[r][c].Flag = 1
}

func (b *Board7x7) Print() {
	for r := range b {
		for c := range b[r] {
			if c > 0 {
				fmt.Print("  ")
			}
			fmt.Fprint(color.Output, b[r][c].Text)
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", 20))
}
