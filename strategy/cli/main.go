package main

import (
	"flag"
	"log"
	"os"

	"github.com/pilagod/go-design-patterns/strategy/shapes"
)

var output = flag.String("output", "text", "The output to use between 'text' and 'image'")

func main() {
	flag.Parse()

	activeStrategy, err := shapes.NewPrinter(*output)

	if err != nil {
		log.Fatal(err)
	}
	switch *output {
	case shapes.TextStrategy:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.ImageStrategy:
		file, err := os.Create("./image.jpg")

		if err != nil {
			log.Fatal("Error opening image")
		}
		defer file.Close()
		activeStrategy.SetWriter(file)
	}
	err = activeStrategy.Print()

	if err != nil {
		log.Fatal(err)
	}
}
