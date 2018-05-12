package shapes

import (
	"fmt"
	"os"

	"github.com/pilagod/go-design-patterns/strategy"
)

const (
	TextStrategy  = "text"
	ImageStrategy = "image"
)

func NewPrinter(strategyType string) (strategy.PrintStrategy, error) {
	switch strategyType {
	case TextStrategy:
		return &TextSquare{
			PrintOutput: strategy.PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case ImageStrategy:
		return &ImageSquare{
			PrintOutput: strategy.PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy %s not found", strategyType)
	}
}
