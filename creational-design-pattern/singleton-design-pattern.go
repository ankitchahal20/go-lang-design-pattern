package singleton

/*

Singleton Design Pattern makes sure that a unique instance of a type is created in the entire program.

Singleton is a creational design pattern, which ensures that only one object of its kind exists
and provides a single point of access to it for any other code.

*/

/*

1. As the name implies, it will provide you with a single instance of an object, and guarantee that there are no
duplicates.
2. At the first call to use the instance, it is created and then reused between all the parts in the application
that need to use that particular behavior.

Example : Singleton pattern in many different situations.
ex 1 : When you want to use the same connection to a database to make every query
ex 2 : When you open a Secure Shell (SSH) connection to a server to do a few tasks,
and don't want to reopen the connection for each task


Objectives
As a general guide, we consider using the Singleton pattern when the following rule applies:
We need a single, shared value, of some particular type.
We need to restrict object creation of some type to a single unit along the entire program.

*/

// First, we create a structure that contains the object which we want to guarantee to be a Singleton
// during the execution of the program:
type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

// Similar example are avialable at https://refactoring.guru/design-patterns/singleton/go/example#example-0
