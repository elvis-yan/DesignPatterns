package main

import (
	"container/list"
	"fmt"
	"reflect"
)

type Subject interface {
	Register(Observer)
	Remove(Observer)
	NotifyAll()
}

type Observer interface {
	Update(temp, humidity, pressure float32)
}

type Displayer interface {
	Display()
}

type WeatherData struct {
	observers   *list.List
	temperature float32
	humidity    float32
	pressure    float32
}

func NewWeatherData() *WeatherData {
	return &WeatherData{observers: list.New()}
}

func (w *WeatherData) Register(o Observer) {
	w.observers.PushBack(o)
}

func (w *WeatherData) Remove(o Observer) {
	elem := w.observers.Front()
	for elem != nil {
		if reflect.DeepEqual(elem.Value, o) {
			w.observers.Remove(elem)
			return
		}
		elem = elem.Next()
	}
}

func (w *WeatherData) NotifyAll() {
	elem := w.observers.Front()
	for elem != nil {
		elem.Value.(Observer).Update(w.temperature, w.humidity, w.pressure)
		elem = elem.Next()
	}
}

func (w *WeatherData) changed() {
	w.NotifyAll()
}

func (w *WeatherData) SetMeasurements(temperature, humidity, pressure float32) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.changed()
}

type CurrentCondition struct {
	temperature, humidity float32
	weatherData           Subject
}

func NewCurrentCondition(weatherData Subject) *CurrentCondition {
	cond := &CurrentCondition{weatherData: weatherData}
	weatherData.Register(cond)
	return cond
}

func (c *CurrentCondition) Update(temperature, humidity, pressure float32) {
	c.temperature = temperature
	c.humidity = humidity
	c.Display()
}

func (c *CurrentCondition) Display() {
	fmt.Printf("Current conditions: %.2f F degrees and %.2f%% humidity\n", c.temperature, c.humidity)
}

type StatisticsCondition struct {
	temperature, humidity float32
	weatherData           Subject
}

func NewStatisticsCondition(weatherData Subject) *StatisticsCondition {
	cond := &StatisticsCondition{weatherData: weatherData}
	weatherData.Register(cond)
	return cond
}

func (s *StatisticsCondition) Update(temperature, humidity, pressure float32) {
	s.temperature = temperature
	s.humidity = humidity
	s.Display()
}

func (s *StatisticsCondition) Display() {
	fmt.Println("Statistics conditions: 🤡")
}

type ForecastCondition struct {
	temperature, humidity float32
	weatherData           Subject
}

func NewForecastCondition(weatherData Subject) *ForecastCondition {
	cond := &ForecastCondition{weatherData: weatherData}
	weatherData.Register(cond)
	return cond
}

func (f *ForecastCondition) Update(temperature, humidity, pressure float32) {
	f.temperature = temperature
	f.humidity = humidity
	f.Display()
}

func (f *ForecastCondition) Display() {
	fmt.Println("Forecast conditions: 😈")
}

func main() {
	weatherData := NewWeatherData()

	curCond := NewCurrentCondition(weatherData)
	statCond := NewStatisticsCondition(weatherData)
	foreCond := NewForecastCondition(weatherData)

	_ = curCond
	_ = statCond
	_ = foreCond

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)

	fmt.Println("================== remove statCond")
	weatherData.Remove(statCond)
	weatherData.SetMeasurements(80, 65, 30.4)

	fmt.Println("==================remove foreCond")
	weatherData.Remove(foreCond)
	weatherData.SetMeasurements(80, 65, 30.4)

	fmt.Println("==================remove curCond")
	weatherData.Remove(curCond)
	weatherData.SetMeasurements(80, 65, 30.4)

}
