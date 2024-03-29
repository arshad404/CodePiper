package main

import (
	"sync"
)

type singleton struct {
	name string
}

// constructor
func newSingletonInstance(input string) *singleton {
	return &singleton{
		name: input,
	}
}

var instance *singleton
var once sync.Once

// GetInstance returns the singleton instance
func GetInstance(name string) *singleton {
	// Called only once
	once.Do(func() {
		instance = newSingletonInstance(name)
		GetInstance("xyz")
	})
	return instance
}

func main2() {
	instance1 := GetInstance("Subscribe to CodePiper Instance 1")
	instance2 := GetInstance("Subscribe to CodePiper Instance 2")
	instance3 := GetInstance("Subscribe to CodePiper Instance 3")

	println("This is instance1 name", instance1.name)
	println("This is instance2 name", instance2.name)
	println("This is instance2 name", instance3.name)
}
