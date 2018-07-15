package main

import (
	"fmt"
)

type Duck interface {
	Display()
	Swim()
	Flyer
	Quacker

	SetFlyer(f Flyer)
	SetQuacker(q Quacker)
}

//__________________________________ Swim __________________________________

type Swimmer struct{}

func (Swimmer) Swim() { fmt.Println("All ducks float, even decoys") }

//__________________________________ Fly __________________________________

type Flyer interface{ Fly() }

type FlyWithWings struct{}

func (*FlyWithWings) Fly() { fmt.Println("I'm flying!!") }

type FlyNoWay struct{}

func (*FlyNoWay) Fly() { fmt.Println("I can't fly") }

type FlyRocketPowered struct{}

func (*FlyRocketPowered) Fly() { fmt.Println("I'm flying with a rocket!") }

//__________________________________ Quack __________________________________

type Quacker interface{ Quack() }

type Quack struct{}

func (*Quack) Quack() { fmt.Println("Quack") }

type MuteQuack struct{}

func (*MuteQuack) Quack() { fmt.Println("<< Silence >>") }

type Squeak struct{}

func (*Squeak) Quack() { fmt.Println("Squeak") }

//__________________________________ Specific Duck __________________________________

type MallardDuck struct {
	Swimmer
	mallardAttr interface{}
	Quacker     Quacker
	Flyer       Flyer
}

func NewMallardDuck(attr interface{}, q Quacker, f Flyer) *MallardDuck {
	return &MallardDuck{
		mallardAttr: attr,
		Quacker:     q,
		Flyer:       f}
}

func (d *MallardDuck) SetFlyer(f Flyer)     { d.Flyer = f }
func (d *MallardDuck) SetQuacker(q Quacker) { d.Quacker = q }

func (d *MallardDuck) Quack()   { d.Quacker.Quack() }
func (d *MallardDuck) Fly()     { d.Flyer.Fly() }
func (d *MallardDuck) Display() { fmt.Println("I'm a real mallard duck   😀") }

type ModelDuck struct {
	Swimmer
	mallardAttr interface{}
	Quacker     Quacker
	Flyer       Flyer
}

func NewModelDuck(attr interface{}, q Quacker, f Flyer) *ModelDuck {
	return &ModelDuck{
		mallardAttr: attr,
		Quacker:     q,
		Flyer:       f}
}

func (d *ModelDuck) SetFlyer(f Flyer)     { d.Flyer = f }
func (d *ModelDuck) SetQuacker(q Quacker) { d.Quacker = q }

func (d *ModelDuck) Quack()   { d.Quacker.Quack() }
func (d *ModelDuck) Fly()     { d.Flyer.Fly() }
func (d *ModelDuck) Display() { fmt.Println("I'm a model duck   😈") }

func main() {
	var duck Duck
	duck = NewMallardDuck("an-attr", new(Quack), new(FlyWithWings))
	duck.Display()
	duck.Swim()
	duck.Quack()
	duck.Fly()

	duck = NewModelDuck("another-attr", new(Quack), new(FlyWithWings))
	duck.Display()
	duck.Swim()
	duck.Quack()
	duck.Fly()

	duck.SetQuacker(new(Squeak))
	duck.SetFlyer(new(FlyRocketPowered))
	duck.Swim()
	duck.Quack()
	duck.Fly()
}
