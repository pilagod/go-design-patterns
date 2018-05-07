package prototype

import (
	"fmt"
)

const (
	// White is color of shirt
	White ShirtColor = 1
	// Black is color of shirt
	Black ShirtColor = 2
	// Blue is color of shirt
	Blue ShirtColor = 3
)

// ItemInfoGetter is an interface that can get item info
type ItemInfoGetter interface {
	GetInfo() string
}

// ShirtColor is type of color of shirt
type ShirtColor byte

// Shirt is item of shirt, implements ItemInfoGetter
type Shirt struct {
	Color ShirtColor
	Price float32
	SKU   string
}

// GetInfo returns info of Shirt
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

var whitePrototype = &Shirt{
	Color: White,
	Price: 15.00,
	SKU:   "empty",
}

var blackPrototype = &Shirt{
	Color: Black,
	Price: 16.00,
	SKU:   "empty",
}

var bluePrototype = &Shirt{
	Color: Blue,
	Price: 17.00,
	SKU:   "empty",
}
