package main

import (
	"MonteCarloTreeSearchInGo/engine"
	"fmt"
)

func main() {
	//idx := 0
	var game engine.Engine
	var location engine.BoardLocation
	game.InitializeGame()
	location = engine.BoardLocation{X: 2, Y: 2}
	_ = game.PlayMove(1, location)
	location = engine.BoardLocation{X: 0, Y: 0}
	_ = game.PlayMove(-1, location)
	location = engine.BoardLocation{X: 1, Y: 1}
	_ = game.PlayMove(1, location)
	location = engine.BoardLocation{X: 1, Y: 0}
	_ = game.PlayMove(-1, location)
	location = engine.BoardLocation{X: 1, Y: 2}
	_ = game.PlayMove(1, location)
	location = engine.BoardLocation{X: 2, Y: 0}
	_ = game.PlayMove(-1, location)
	fmt.Println(game.CheckEnd())
}
