package engine

import "fmt"

// Engine class defines the game state and resolve moves
type Engine struct {
	state     [][]int32
	whoToPlay int32 // 1 or -1
}

// BoardLocation class
type BoardLocation struct {
	X uint16
	Y uint16
}

// Idx2BL select a board location using an index number from 0 to 8
func (b *BoardLocation) Idx2BL(idx int32) BoardLocation {

	return BoardLocation{uint16(idx % 3), uint16((idx - idx%3) / 3)}
}

// InitializeGame initialize the game
func (e *Engine) InitializeGame() {
	e.state = [][]int32{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	e.whoToPlay = 1
}

// PlayMove play a move
func (e *Engine) PlayMove(player int32, loc BoardLocation) bool {
	if e.state[loc.X][loc.Y] == 0 && e.whoToPlay == player && (player == 1 || player == -1) {
		e.state[loc.X][loc.Y] = player
		e.whoToPlay = -player
		e.PrintState()
		return false
	}
	return true
}

// PrintState print the state of the game
func (e *Engine) PrintState() {

	for _, a := range e.state {
		fmt.Println(a)
	}
	fmt.Println()
}

// CheckEnd check if the game has ended. Returns a bool true if the game ended and false if not, and a int32 that represents the player that won the game
func (e *Engine) CheckEnd() (bool, int32) {

	// check rows
	for _, a := range e.state {
		if isEqual(a...) == true {
			return true, a[0]
		}
	}

	var aux []int32 = []int32{0, 0, 0}

	// check diagonals
	for i := 0; i < len(e.state); i++ {
		aux[i] = e.state[i][i]
	}
	if isEqual(aux...) == true {
		return true, aux[0]
	}
	for i := 0; i < len(e.state); i++ {
		aux[i] = e.state[len(e.state)-i-1][i]
	}
	if isEqual(aux...) == true {
		return true, aux[0]
	}
	// check columns
	for i := 0; i < len(e.state); i++ {
		for j := 0; j < len(e.state); j++ {
			aux[j] = e.state[j][i]
		}
		if isEqual(aux...) == true {
			return true, aux[0]
		}
	}

	return false, 0
}

func isEqual(nums ...int32) bool {
	aux := make(map[int32]bool)
	for _, num := range nums {
		aux[num] = false
	}
	if len(aux) == 1 {
		return true
	}
	return false
}
