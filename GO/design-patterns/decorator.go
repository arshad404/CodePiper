package main

import "fmt"

// STEP-1
// Coffee represents the base component interface
type Coffee interface {
	GetDescription() string
	Cost() float64
}

// STEP-2
// SimpleCoffee represents the concrete component
type SimpleCoffee struct{}

func (c *SimpleCoffee) GetDescription() string {
	return "Simple Coffee"
}

func (c *SimpleCoffee) Cost() float64 {
	return 1.0
}

// STEP-3
// CoffeeDecorator represents the decorator abstract class
type CoffeeDecorator struct {
	CoffeeComponent Coffee
}

func (cd *CoffeeDecorator) GetDescription() string {
	return cd.CoffeeComponent.GetDescription()
}

func (cd *CoffeeDecorator) Cost() float64 {
	return cd.CoffeeComponent.Cost()
}

// Milk represents a concrete decorator
type Milk struct {
	*CoffeeDecorator
}

func NewMilk(coffee Coffee) *Milk {
	return &Milk{&CoffeeDecorator{coffee}}
}

func (m *Milk) GetDescription() string {
	return m.CoffeeDecorator.GetDescription() + ", Milk"
}

func (m *Milk) Cost() float64 {
	return m.CoffeeDecorator.Cost() + 0.5
}

// Sugar represents another concrete decorator
type Sugar struct {
	*CoffeeDecorator
}

func NewSugar(coffee Coffee) *Sugar {
	return &Sugar{&CoffeeDecorator{coffee}}
}

func (s *Sugar) GetDescription() string {
	return s.CoffeeDecorator.GetDescription() + ", Sugar"
}

func (s *Sugar) Cost() float64 {
	return s.CoffeeDecorator.Cost() + 0.25
}

func main() {
	// Create a simple coffee
	coffee := &SimpleCoffee{}

	// Add milk to the coffee
	coffeeWithMilk := NewMilk(coffee)

	// Add sugar to the coffee
	coffeeWithMilkAndSugar := NewSugar(coffeeWithMilk)

	// Print description and cost of the decorated coffee
	fmt.Println("Description:", coffeeWithMilkAndSugar.GetDescription())
	fmt.Println("Cost:", coffeeWithMilkAndSugar.Cost())
}
