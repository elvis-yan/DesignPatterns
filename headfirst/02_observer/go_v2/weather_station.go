package main

import (
	"container/list"
	"fmt"
	"reflect"
	"sync"
)

//______________________________________ interfaces _______________________________________

type Subject interface {
	Register(Observer)
	Remove(Observer)
	NotifyAll()
}

type Observer interface {
	Update()
	Recv(Data)
}

// type Displayer interface {
// 	Display()
// }

//______________________________________ Weather - Subject _______________________________________

type Data struct {
	Temperature, Humidity, Pressure float32
}

type Weather struct {
	eventCh chan Data
	data    Data

	wg sync.WaitGroup

	mu        sync.RWMutex
	observers *list.List
}

func Newweather() *Weather {
	w := &Weather{eventCh: make(chan Data, 1), observers: list.New()}
	go w.NotifyAll()
	return w
}

func (w *Weather) Register(o Observer) {
	w.observers.PushBack(o)

}

func (w *Weather) Remove(o Observer) {
	elem := w.observers.Front()
	for elem != nil {
		if reflect.DeepEqual(elem.Value, o) {
			w.observers.Remove(elem)
			return
		}
		elem = elem.Next()
	}
}

func (w *Weather) NotifyAll() {
	for {

		data := <-w.eventCh
		elem := w.observers.Front()
		for elem != nil {
			w.wg.Add(1) // move up will panic: sync: negative WaitGroup counter :107
			elem.Value.(Observer).Recv(data)
			elem = elem.Next()
		}

	}
}

func (w *Weather) SetMeasurements(data Data) {
	w.data = data
	w.eventCh <- data
}

//__________________________________ Condition - Observer ___________________________________

type Condition struct {
	Ch      chan Data
	data    Data
	weather Subject
	Display func()
}

func NewCondition(weather Subject) *Condition {
	cond := &Condition{Ch: make(chan Data), weather: weather}
	cond.Display = cond.DisplayCurrent
	weather.Register(cond)
	go cond.Update()
	return cond
}

func (c *Condition) Update() {
	for {
		c.data = <-c.Ch
		c.Display()
		c.weather.(*Weather).wg.Done()
	}
}

func (c *Condition) Recv(data Data) {
	c.Ch <- data
}

func (c *Condition) DisplayCurrent() {
	fmt.Println("Current conditions: 😈")
}

func (c *Condition) DisplayStatistics() {
	fmt.Println("Statistics conditions: 🤡")
}

func (c *Condition) DisplayForecast() {
	fmt.Println("Forecast conditions: 😀")
}

func main() {
	weather := Newweather()

	curCond := NewCondition(weather)
	curCond.Display = curCond.DisplayCurrent
	foreCond := NewCondition(weather)
	foreCond.Display = foreCond.DisplayForecast
	statCond := NewCondition(weather)
	statCond.Display = statCond.DisplayStatistics

	weather.SetMeasurements(Data{80, 65, 30.4})
	weather.SetMeasurements(Data{82, 70, 29.2})
	weather.SetMeasurements(Data{78, 90, 29.2})

	weather.wg.Wait()
}
