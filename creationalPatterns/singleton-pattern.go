package creationalPatterns

/*
The Singleton Pattern is a creational design pattern that ensures a class has only one instance,
and provides a global point of access to that instance.
It is useful when you want to limit the number of instances of a class to one, to ensure that there is only one point of
control or coordination for a particular resource or functionality.
*/

/*
In this example, we define a Singleton struct, which has a data attribute. We also define a global variable instance of type *Singleton, which will hold the singleton instance.

We define a GetInstance() function, which returns the singleton instance. If the instance does not yet exist, it is created.

We define methods SetData() and GetData() for the Singleton struct, which allow us to set and get the data attribute of the singleton instance.

Finally, we demonstrate the use of the Singleton Pattern by getting the singleton instance, setting its data to "new data", and then getting the data back out.

Overall, the Singleton Pattern ensures that a class has only one instance, and provides a global point of access to that instance. This can be useful in scenarios where you need to limit the number of instances of a class to one, to ensure that there is only one point of control or coordination for a particular resource or functionality.
*/

// Singleton Class

type Singleton struct {
	data string
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{data: "initial data"}
	}
	return instance
}

func (s *Singleton) SetData(data string) {
	s.data = data
}

func (s *Singleton) GetData() string {
	return s.data
}
