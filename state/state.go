package state

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type GameState interface {
	ExecuteState(*GameContext) bool
}

type GameContext struct {
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}

type StartState struct{}

func (ss *StartState) ExecuteState(gc *GameContext) bool {
	rand.Seed(time.Now().UnixNano())

	gc.SecretNumber = rand.Intn(10)

	fmt.Println("Introduce a number of retries to set the difficulty:")
	fmt.Fscanf(os.Stdin, "%d", &gc.Retries)

	gc.Next = &AskState{}

	return true
}

type AskState struct{}

func (as *AskState) ExecuteState(gc *GameContext) bool {
	var n int

	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", gc.Retries)
	fmt.Fscanf(os.Stdin, "%d", &n)

	gc.Retries = gc.Retries - 1

	if n == gc.SecretNumber {
		gc.Won = true
		gc.Next = &FinishState{}
	} else if gc.Retries == 0 {
		gc.Next = &FinishState{}
	}
	return true
}

type FinishState struct{}

func (fs *FinishState) ExecuteState(gc *GameContext) bool {
	if gc.Won {
		gc.Next = &WinState{}
	} else {
		gc.Next = &LoseState{}
	}
	return true
}

type WinState struct{}

func (ws *WinState) ExecuteState(gc *GameContext) bool {
	fmt.Printf("Congrats, you won! Secret number is %d\n", gc.SecretNumber)
	return false
}

type LoseState struct{}

func (ls *LoseState) ExecuteState(gc *GameContext) bool {
	fmt.Printf("Sorry, you lost. Secret number is %d\n", gc.SecretNumber)
	return false
}
