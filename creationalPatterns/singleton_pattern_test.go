package creationalPatterns

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	// Test that GetInstance returns the same instance every time
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Errorf("Expected instance1 to equal instance2 but got %v and %v", instance1, instance2)
	}
}

func TestConcurrency(t *testing.T) {
	instance := GetInstance()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		instance.IncrementCount()
	}()

	go func() {
		defer wg.Done()
		instance.IncrementCount()
	}()

	wg.Wait()

	if instance.GetCount() != 2 {
		t.Errorf("Expected count to be 2, but got %d instead", instance.GetCount())
	}
}
