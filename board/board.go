package board

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
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

func (b *Board7x7) CanSet(x int, y int) (ok bool) {
	if x > 6 || y > 6 {
		return false
	}
	return b[x][y].Flag == 0
}
func (b *Board7x7) Set(Text string, x int, y int) {
	b[x][y].Text = Text
	b[x][y].Flag = 1
}

func (b *Board7x7) Print() {
	for i := range b {
		for j := range b[i] {
			if j > 0 {
				fmt.Print("  ")
			}
			fmt.Fprint(color.Output, b[i][j].Text)
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", 20))
}
