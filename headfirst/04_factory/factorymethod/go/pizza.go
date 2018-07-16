package main

import (
	"fmt"
)

const (
	// Pizza Types
	Cheese PizzaType = "Cheese Pizza"

	// Store Style
	NY      Style = "NY"
	Chicago Style = "Chicago"
)

type PizzaType string
type Style string

type Pizza struct {
	Style    Style
	Type     PizzaType
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

var PizzaTable = map[Style]map[PizzaType]*Pizza{
	NY:      {Cheese: newChicagoStyleCheesePizza()},
	Chicago: {Cheese: newChicagoStyleCheesePizza()},
}

func newChicagoStyleCheesePizza() *Pizza {
	pizza := &Pizza{}
	pizza.Name = "Chicago Style Deep Dish Cheese Pizza"
	pizza.Dough = "Extra Thick Crust Dough"
	pizza.Sauce = "Plum Tomato Sauce"
	pizza.Toppings = []string{"Shredded Mozzarella Cheese"}
	return pizza
}

func newNYStyleCheesePizza() *Pizza {
	pizza := &Pizza{}
	pizza.Name = "NY Style Sauce and Cheese Pizza"
	pizza.Dough = "Thin Crust Dough"
	pizza.Sauce = "Marinara Sauce"
	pizza.Toppings = []string{"Grated Reggiano Cheese"}
	return pizza
}

func NewPizza(style Style, typ PizzaType) *Pizza {
	return PizzaTable[style][typ]
}

//_________________________________________________________________

type Store struct {
	Style Style
}

func NewStore(style Style) *Store {
	return &Store{Style: style}
}

func (s *Store) Order(typ PizzaType) *Pizza {
	var pizza *Pizza
	switch typ {
	case Cheese:
		pizza = NewPizza(s.Style, typ)
	}

	s.Prepare(*pizza)
	s.Bake()
	s.Cut()
	s.Box()

	return pizza
}

func (s *Store) Prepare(p Pizza) {
	fmt.Println("Preparing", p.Name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings:")
	for _, t := range p.Toppings {
		fmt.Println("  ", t)
	}
}

func (s *Store) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (s *Store) Cut() {
	switch s.Style {
	case NY:
		fmt.Println("Cutting the pizza into diagonal slices")
	case Chicago:
		fmt.Println("Cutting the pizza into square slices")
	}
}
func (s *Store) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func main() {
	nyStore := NewStore(NY)
	chicagoStore := NewStore(Chicago)

	pizza := nyStore.Order(Cheese)
	fmt.Printf("Ethan ordered a %s\n\n", pizza.Name)

	pizza = chicagoStore.Order(Cheese)
	fmt.Printf("Joel ordered a %s\n\n", pizza.Name)
}
