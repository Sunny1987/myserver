package handlers

import (
	"encoding/json"
	"myserver/myserver/data"
	"net/http"
)

// CreatePOSTForPersons method helps to add a person in the database
func (p *Persons) CreatePOSTForPersons(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("\"************CreatePOSTForPersons API ******************\"")

	//Update the response header for JSON
	rw.Header().Add("content-type", "application/json")

	//get the database instance
	collection, ctx := data.GetDB(p.l, "persons")

	//get the person instance context post middleware
	person := r.Context().Value(KeyPerson{}).(*data.Person)

	//perform DB insertion
	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		p.l.Printf("Error in inserting data: %v", err)
		http.Error(rw, "Failed to updated DB", http.StatusInternalServerError)
	}
	p.l.Printf("The item is inserted in DB: %v", result)

	//Write output response
	json.NewEncoder(rw).Encode(result)

}

// CreatePOSTForProducts method helps to add a person in the database
func (pr *Products) CreatePOSTForProducts(rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("\"************CreatePOSTForProducts API ******************\"")

	//Update the response header for JSON
	rw.Header().Add("content-type", "application/json")

	//get the database instance
	collection, ctx := data.GetDB(pr.l, "products")

	//get the person instance context post middleware
	product := r.Context().Value(KeyPerson{}).(*data.Product)

	//perform DB insertion
	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		pr.l.Printf("Error in inserting data: %v", err)
		http.Error(rw, "Failed to updated DB", http.StatusInternalServerError)
	}
	pr.l.Printf("The item is inserted in DB: %v", result)

	//Write output response
	json.NewEncoder(rw).Encode(result)

}
