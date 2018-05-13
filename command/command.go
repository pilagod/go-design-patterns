package command

import (
	"fmt"
	"time"
)

type Command interface {
	Execute()
}

type ConsoleOutputCommand struct {
	message string
}

func (coc *ConsoleOutputCommand) Execute() {
	fmt.Println(coc.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")
	return &ConsoleOutputCommand{message: s}
}

type CommandQueue struct {
	queue []Command
}

func (cq *CommandQueue) AddCommand(c Command) {
	cq.queue = append(cq.queue, c)

	if len(cq.queue) == 3 {
		for _, command := range cq.queue {
			command.Execute()
		}
		cq.queue = make([]Command, 3)
	}
}

type InfoCommand interface {
	Info() string
}

type TimeInfoCommand struct {
	Start time.Time
}

func (tic *TimeInfoCommand) Info() string {
	return time.Since(tic.Start).String()
}

type HelloMessageInfoCommand struct {
	Message string
}

func (hmic *HelloMessageInfoCommand) Info() string {
	return "Hello world!"
}
