package builder

/*

The Builder pattern helps us construct complex objects without directly instantiating their struct,
or writing the logic they require.

Imagine an object that could have dozens of fields that are more complex structs themselves.
Now imagine that you have many objects with these characteristics, and you could have more.
We don't want to write the logic to create all these objects in the package that just needs to use the objects.

*/

/*
A Builder design pattern tries to:
Abstract complex creations so that object creation is separated from the object user
Create an object step by step by filling its fields and creating the embedded objects
Reuse the object creation algorithm between many objects


Separate the construction of a complex object from its representation.
Allow the same construction process to create different representations.
Encapsulate the construction logic in separate builder classes.
Enable the creation of objects with many optional or varying parts.

*/

/*

The builder pattern is a creational design pattern that separates the construction of a complex object from its representation.
It allows for the creation of different representations of the object using the same construction process.
The pattern involves a Builder interface, Concrete Builders, a Director, and the Product.

Key components of the builder pattern:

Product: The complex object being constructed.
Builder interface: Specifies methods for constructing parts of the complex object.
Concrete Builders: Implement the Builder interface, providing specific implementations for each part.
Director: Constructs the complex object using the Builder interface, orchestrating the construction process step by step.

*/

// Step 1: Define the Product with the required properties.
type Product struct {
	Wheels    int
	Seats     int
	Structure string
}

// ^^ The product represents the complex object that we want to build, it is always defined first

// Step 2: Create a Builder interface with methods to construct the object
type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	GetVehicle() Product
}

// define the Builder interface. Builder is the interface that specifies the methods for building the product parts.
// CarBuilder, BikeBuilder and ShuttleBusBuilder are the implementations of the Builder interface

// CarBuilder is used to concrete Builder structs, it is step 3
type CarBuilder struct {
	v Product
}

func (c *CarBuilder) SetWheels() Builder {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() Builder {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() Builder {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() Product {
	return c.v
}

// BikeBuilder is used to concrete Builder structs
type BikeBuilder struct {
	v Product
}

func (b *BikeBuilder) SetWheels() Builder {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() Builder {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() Builder {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() Product {
	return b.v
}

// ShuttleBusBuilder is used to concrete Builder structs
type ShuttleBusBuilder struct {
	v Product
}

func (s *ShuttleBusBuilder) SetWheels() Builder {
	s.v.Wheels = 4 * 2
	return s
}

func (s *ShuttleBusBuilder) SetSeats() Builder {
	s.v.Seats = 30
	return s
}

func (s *ShuttleBusBuilder) SetStructure() Builder {
	s.v.Structure = "ShuttleBus"
	return s
}

func (s *ShuttleBusBuilder) GetVehicle() Product {
	return s.v
}

// ManufactureDirector is used to take a Builder interface and construct the object.
type ManufactureDirector struct {
	builder Builder
}

// ManufactureDirector is responsible for using the Builder interface to make the product.

func (f *ManufactureDirector) SetBuilder(b Builder) {
	f.builder = b
}

func (f *ManufactureDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}
