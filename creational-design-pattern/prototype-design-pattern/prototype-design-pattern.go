package prototype

import (
	"errors"
	"fmt"
)

/*
While with the Builder pattern, we are dealing with repetitive building algorithms and
with the factories we are simplifying the creation of many types of objects;
with the Prototype pattern, we will use an already created instance of some type to clone it
and complete it with the particular needs of each context.
*/

// Description : The aim of the Prototype pattern is to have an object or a set of objects that is already created at compilation time,
// but which you can clone as many times as you want at runtime.

const (
	White = 1
	Black = 2
	Blue  = 3
)

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

func GetShirtsCloner() ShirtCloner {
	return nil
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

func (i *Shirt) GetPrice() float32 {
	return i.Price
}

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}
