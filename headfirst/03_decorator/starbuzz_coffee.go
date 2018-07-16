package main

import (
	"fmt"
)

const (
	Mocha CondimentType = "Mocha"
	Soy   CondimentType = "Soy"
	Whip  CondimentType = "Whip"
)

const (
	Espresso   CoffeeType = "Espresso"
	HouseBlend CoffeeType = "House Blend Coffee"
)

var CondimentPriceTable = map[CondimentType]float32{
	Mocha: .20,
	Soy:   .15,
	Whip:  .10,
}

var CoffeePriceTable = map[CoffeeType]float32{
	Espresso:   1.99,
	HouseBlend: .89,
}

type CondimentType string
type CoffeeType string

// ____________________ Condiment _________________________

type Condiment struct {
	Type  CondimentType
	Price float32
}

func NewCondiment(typ CondimentType) *Condiment {
	return &Condiment{
		Type:  typ,
		Price: CondimentPriceTable[typ],
	}
}

//_____________________ Some Kind of Coffee ____________________

type Coffee struct {
	Type       CoffeeType
	Price      float32
	Condiments []*Condiment
}

func NewCoffee(typ CoffeeType) *Coffee {
	return &Coffee{
		Type:       typ,
		Price:      CoffeePriceTable[typ],
		Condiments: make([]*Condiment, 0),
	}
}

func (c *Coffee) Add(cdm *Condiment) {
	c.Condiments = append(c.Condiments, cdm)
}

func (c *Coffee) Cost() float32 {
	var price float32
	for _, cdm := range c.Condiments {
		price += cdm.Price
	}
	return price + c.Price
}

func (c *Coffee) GetDescription() string {
	var decorates string
	for _, cdm := range c.Condiments {
		decorates += string(cdm.Type) + ", "
	}
	return decorates + string(c.Type)
}

func main() {
	beverage := NewCoffee(Espresso)
	fmt.Printf("Desc: %s, Price: %.2f $\n\n", beverage.GetDescription(), beverage.Cost())

	beverage2 := NewCoffee(HouseBlend)
	beverage2.Add(NewCondiment(Mocha))
	beverage2.Add(NewCondiment(Mocha))
	beverage2.Add(NewCondiment(Whip))
	fmt.Printf("Desc: %s, Price: %.2f $\n\n", beverage2.GetDescription(), beverage2.Cost())

	beverage3 := NewCoffee(Espresso)
	beverage3.Add(NewCondiment(Soy))
	beverage3.Add(NewCondiment(Mocha))
	beverage3.Add(NewCondiment(Whip))
	fmt.Printf("Desc: %s, Price: %.2f $\n\n", beverage3.GetDescription(), beverage3.Cost())
}
