package factory

/*

	The Factory method pattern (or simply, Factory) is probably the second-best known and used design pattern in the industry.

	Its purpose is to abstract the user from the knowledge of the struct he needs to achieve for a specific purpose,
	such as retrieving some value, maybe from a web service or a database.

	Factory can provide an interface that fits the user needs.
	It also eases the process of downgrading or upgrading of the implementation of the underlying type if needed.

	Imagine that you want to organize your holidays using a trip agency.
	You don't deal with hotels and traveling and you just tell the agency the destination you are interested in so that they provide you with everything you need.
	The trip agency represents a Factory of trips.
*/

/*
Objectives
	After the previous description, the following objectives of the Factory Method design pattern must be clear to you:

	Delegating the creation of new instances of structures to a different part of the program

	Working at the interface level instead of with concrete implementations Grouping families of objects to obtain a family object creator
*/

/*
Example:
	For our example, we are going to implement a payments method Factory, which is going to provide us with different ways of paying at a shop.
	In the beginning, we will have two methods of paying i.e. cash and credit card.
	We'll also have an interface with the method, Pay , which every struct that wants to be used as a payment method must implement.
*/

/*
Acceptance criteria

	Using the previous description, the requirements for the acceptance criteria are the following:
	To have a common method for every payment method called Pay
	To be able to delegate the creation of payments methods to the Factory
	To be able to add more payment methods to the library by just adding it to the factory method

*/

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash       = 1
	DebitCard  = 2
	CreditCard = 3
)

type CashPM struct{}
type DebitCardPM struct{}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {

	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCard:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

// upgrading
type CreditCardPM struct{}

func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using new credit card implementation\n", amount)
}
