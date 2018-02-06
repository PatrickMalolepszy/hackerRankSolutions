package main

import (
	"fmt"
)

type bomberman struct {
	r, c  int // dimensions of the game board.
	board board
	temp  board
}

type board [][]int

func newBomberman(r, c int, startState board) *bomberman {
	return &bomberman{
		board: startState,
		temp:  newBoard(r, c),
		r:     r,
		c:     c,
	}
}

func (b *bomberman) detonate() {
	prev := b.board
	next := b.temp
	next.copy(prev)
	// find any bombs about to go off and detonate.
	for i := range prev {
		for j := range prev[i] {
			if prev[i][j] == 0 { // bomb going off
				next.detonateBomb(i, j, b.r, b.c)
			}
		}
	}
	b.board = next
}

func (b *bomberman) fill() {
	cur := b.board
	// fill any empty spots with new bombs:
	for i := range cur {
		for j := range cur[i] {
			if cur[i][j] == -1 {
				cur[i][j] = 3
			}
		}
	}
}

func (b *bomberman) tick() {
	cur := b.board
	for i := range cur {
		for j := range cur[i] {
			if cur[i][j] >= 0 {
				cur[i][j]--
			}
		}
	}
}

func (b *bomberman) play(ticks int) board {
	detonate := false
	cycleCheck1 := newBoard(b.r, b.c)
	cycleFound := false
	cycleOdd := true

	// bommerman fills up board with bombs before denotation.
	if (ticks > 2) && ((ticks % 2) == 0) {
		return fullBoard(b.r, b.c)
	}

	//fmt.Println("starting board:")
	//fmt.Println(b.board)
	//fmt.Println("Wait one sec:")
	b.tick()
	//fmt.Println(b.board)
	for i := 1; i < ticks; i++ {
		b.tick()
		//fmt.Println("after tick:")
		//fmt.Println(b.board)
		if detonate {
			b.detonate()
			//fmt.Println("detonating bombs:")
			//fmt.Println(b.board)
			if cycleFound {
				return b.board
			}
			if cycleOdd {
				if cycleCheck1.equals(b.board) {
					if ((ticks-1) % 4) != 0 {
						return b.board
					}
					cycleFound = true
				}
				cycleCheck1.copy(b.board)
			}
			cycleOdd = !cycleOdd
		} else {
			b.fill()
			//fmt.Println("filling in bombs:")
		}
		//fmt.Println("after change at tick:", i)
		//fmt.Println(b.board)
		detonate = !detonate
	}
	return b.board
}

func fullBoard(r, c int) board {
	fullBoard := make([][]int, r)
	for i := range fullBoard {
		fullBoard[i] = make([]int, c)
		for j := range fullBoard[i] {
			fullBoard[i][j] = 1
		}
	}
	return fullBoard
}

func (b board) String() string {
	var result string
	for i := range b {
		for j := range b[i] {
			if b[i][j] >= 0 {
				//result += strconv.Itoa(b[i][j])
				result += "O"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}

func (b board) copy(from board) {
	for i := range b {
		for j := range b[i] {
			b[i][j] = from[i][j]
		}
	}
}

func (b board) equals(b2 board) bool {
	for i := range b {
		for j := range b[i] {
			if b[i][j] != b2[i][j] {
				return false
			}
		}
	}
	return true
}

func (b board) detonateBomb(i, j, r, c int) {
	b[i][j] = -1
	// down
	if ((i - 1) >= 0) && (b[i-1][j] != 0) {
		b[i-1][j] = -1
	}
	// up
	if ((i + 1) < r) && (b[i+1][j] != 0) {
		b[i+1][j] = -1
	}
	// left
	if ((j - 1) >= 0) && (b[i][j-1] != 0) {
		b[i][j-1] = -1
	}
	// right
	if ((j + 1) < c) && (b[i][j+1] != 0) {
		b[i][j+1] = -1
	}
}

func main() {
	var r, c, n int
	var row string
	fmt.Scan(&r, &c, &n)
	startState := newBoard(r, c)
	for i := 0; i < r; i++ {
		fmt.Scan(&row)
		j := 0
		for k := range row {
			if row[k] == 'O' {
				startState[i][j] = 3
				j++
			} else if row[k] == '.' {
				j++
			}
		}
	}
	game := newBomberman(r, c, startState)
	result := game.play(n)
	fmt.Println(result)
}

func newBoard(r, c int) board {
	game := make([][]int, r)
	for i := range game {
		game[i] = make([]int, c)
		for j := range game[i] {
			game[i][j] = -1
		}
	}
	return game
}
