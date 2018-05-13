package main

import (
	"fmt"
	"time"

	"github.com/pilagod/go-design-patterns/command"
)

func main() {
	// basic command
	queue := command.CommandQueue{}

	queue.AddCommand(command.CreateCommand("First message"))
	queue.AddCommand(command.CreateCommand("Second message"))
	queue.AddCommand(command.CreateCommand("Third message"))
	queue.AddCommand(command.CreateCommand("Fourth message"))
	queue.AddCommand(command.CreateCommand("Fifth message"))

	// info command
	timeInfoCommand := &command.TimeInfoCommand{
		Start: time.Now(),
	}
	helloMessageInfoCommand := &command.HelloMessageInfoCommand{}

	time.Sleep(time.Second)

	fmt.Println(timeInfoCommand.Info())
	fmt.Println(helloMessageInfoCommand.Info())
}
