package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"

	"github.com/gonejack/a-puzzle-a-day/board"
	"github.com/gonejack/a-puzzle-a-day/piece"
)

func main() {
	now := time.Now()
	mon := now.Month().String()[:3]
	day := strconv.Itoa(now.Day())

	if len(os.Args) > 2 {
		mon = os.Args[1]
		day = os.Args[2]
	}

	b := board.Board
	for i := range b {
		for j := range b[i] {
			switch {
			case strings.EqualFold(b[i][j].Text, mon):
				b[i][j].Flag = 1
			case strings.EqualFold(b[i][j].Text, day):
				b[i][j].Flag = 1
			}
		}
	}

	fmt.Printf("searching for %s %s\n", mon, day)
	search(&b, 0)
	fmt.Printf("found %d solutions\n", count)
}

var count = 0
var placed = make(map[int]bool)
var cc = []func(string, ...interface{}) string{
	color.RedString,
	color.GreenString,
	color.YellowString,
	color.BlueString,
	color.MagentaString,
	color.CyanString,
	color.HiBlueString,
	color.HiRedString,
}

func search(b *board.Board7x7, idx int) {
	if len(placed) == 8 {
		b.Print()
		count += 1
		return
	}

	row := idx / 7
	col := idx % 7
	if row >= 7 {
		return
	}

	for i := range piece.Pieces {
		if !placed[i] {
			for _, p := range piece.Pieces[i] {
				if p.CanPlace(b, row, col) {
					c := *b

					placed[i] = true
					p.Place(&c, row, col, cc[i]("◼︎"))
					search(&c, idx+1)

					delete(placed, i)
				}
			}
		}
	}

	if !b.CanSet(row, col) {
		search(b, idx+1)
		return
	}
}
