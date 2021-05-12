package handlers

import (
	"log"
)

// Persons This struct is used as object of all handler methods
type Persons struct {
	l *log.Logger

	//dbCol *mongo.Collection
}

// NewPerson This function returns the Persons based handler
func NewPerson(l *log.Logger) *Persons {
	return &Persons{l: l}
}

//Products is used a handler for product handler methods
type Products struct {
	l *log.Logger
}

// NewProduct This function returns the Products based handler
func NewProduct(l *log.Logger) *Products {
	return &Products{l: l}
}

//Transactions is used a handler for transaction handler methods
type Transactions struct {
	l *log.Logger
}

// NewTransaction This function returns the Transactions based handler
func NewTransaction(l *log.Logger) *Transactions {
	return &Transactions{l: l}

}

type Common struct {
	Persons  *Persons
	Products *Products
}

func NewCommon(l *log.Logger) *Common {
	return &Common{
		Persons:  NewPerson(l),
		Products: NewProduct(l),
	}
}
