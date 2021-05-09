package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Person is used as a struct schema for database
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Phone     string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
}
