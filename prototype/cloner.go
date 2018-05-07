package prototype

// ShirtCloner creates basic shirts from given color
type ShirtCloner interface {
	GetClone(color ShirtColor) (ItemInfoGetter, error)
}

// ShirtCache implements ShirtCloner
type ShirtCache struct{}

// GetClone get a new basic shirt of given color
func (s *ShirtCache) GetClone(color ShirtColor) (ItemInfoGetter, error) {
	var newItem Shirt

	switch color {
	case White:
		newItem = *whitePrototype
	case Black:
		newItem = *blackPrototype
	case Blue:
		newItem = *bluePrototype
	}
	return &newItem, nil
}

// GetShirtCloner returns ShirtCloner
func GetShirtCloner() ShirtCloner {
	return &ShirtCache{}
}
