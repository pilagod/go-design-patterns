package main

import "github.com/pilagod/go-design-patterns/state"

func main() {
	game := state.GameContext{
		Next: &state.StartState{},
	}
	for game.Next.ExecuteState(&game) {
	}
}
