package handlers

import (
	"log"
)

// Persons This struct is used as object of  all handler methods
type Persons struct {
	l *log.Logger

	//dbCol *mongo.Collection
}

// NewPerson This function returns the Persons based handler
func NewPerson(l *log.Logger) *Persons {
	return &Persons{l: l}
}
