package main

import (
	"fmt"
	"sync"
)

// Singleton represents the singleton object.
type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

// GetInstance returns the singleton instance.
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Singleton Instance"}
	})
	return instance
}

func main() {
	// Get the singleton instance
	singleton := GetInstance()

	// Print the data
	fmt.Println("Singleton Data:", singleton.data)

	// Output:
	// Singleton Data: Singleton Instance
}
