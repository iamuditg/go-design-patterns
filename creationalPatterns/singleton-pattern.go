package creationalPatterns

/*
The Singleton Pattern is a creational design pattern that ensures a class has only one instance,
and provides a global point of access to that instance.
It is useful when you want to limit the number of instances of a class to one, to ensure that there is only one point of
control or coordination for a particular resource or functionality.
*/

// Singleton Class

type singleton struct {
	count int
}

var instance *singleton

// GetInstance return the singleton instance
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// IncrementCount increments the count variable of the singleton
func (s *singleton) IncrementCount() {
	s.count++
}

// GetCount returns the current value of the count variable
func (s *singleton) GetCount() int {
	return s.count
}
