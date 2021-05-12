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

//Product is a schema for products in database
type Product struct {
	ProdID       primitive.ObjectID `json:"_prod_id,omitempty" bson:"_prod_id,omitempty"`
	ProductName  string             `json:"product_name,omitempty,omitempty" bson:"product_name,omitempty"`
	ProductPrice string             `json:"product_price,omitempty" bson:"product_price,omitempty"`
	Category     string             `json:"category,omitempty" bson:"category,omitempty"`
	ProdQuantity string             `json:"prod_quantity,omitempty" bson:"prod_quantity,omitempty"`
}

// Transaction is a schema for Person in Product interactions in database
type Transaction struct {
	TranID       primitive.ObjectID `json:"tran_id,omitempty" bson:"tran_id,omitempty"`
	PID          primitive.ObjectID `json:"prod_id,omitempty" bson:"prod_id,omitempty"`
	PeID         primitive.ObjectID `json:"pe_id,omitempty" bson:"pe_id,omitempty"`
	FirstName    string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName     string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Phone        string             `json:"phone,omitempty" bson:"phone,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	ProductName  string             `json:"product_name,omitempty,omitempty" bson:"product_name,omitempty"`
	ProductPrice string             `json:"product_price,omitempty" bson:"product_price,omitempty"`
	Quantity     int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
}
