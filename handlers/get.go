package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myserver/myserver/data"
	"net/http"
)

// GetAllPersons This method calls the list of persons in DB
func (p *Persons) GetAllPersons(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("\"************ GetAllPersons API ******************\"")

	//Update the response header for JSON
	rw.Header().Add("content-type", "application/json")

	//get the database instance
	collection, ctx := data.GetDB(p.l, "persons")

	//Extract all records
	cursor, err := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	if err != nil {
		http.Error(rw, "Failed to get data from DB", http.StatusInternalServerError)
		return
	}

	//Collect all records for response
	var persons []data.Person

	//Collect the records and store in slice
	for cursor.Next(ctx) {
		var person data.Person
		cursor.Decode(&person)
		persons = append(persons, person)
	}
	if err := cursor.Err(); err != nil {
		http.Error(rw, "Failed to get data from DB", http.StatusInternalServerError)
		return
	}
	//Write the response
	json.NewEncoder(rw).Encode(persons)
}

// GetOnePerson This method calls one person in DB
func (p *Persons) GetOnePerson(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("\"************ GetOnePerson API ******************\"")

	//Update the response header for JSON
	rw.Header().Add("content-type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
		return
	}

	//get the database instance
	collection, ctx := data.GetDB(p.l, "persons")
	var person data.Person

	//Extract one record
	err = collection.FindOne(ctx, data.Person{ID: id}).Decode(&person)
	if err != nil {
		http.Error(rw, "Unable to parse data", http.StatusInternalServerError)
		return
	}

	//Write the response
	json.NewEncoder(rw).Encode(person)
}

// GetAllProducts This method calls the list of persons in DB
func (pr *Products) GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("\"************ GetAllProducts API ******************\"")

	//Update the response header for JSON
	rw.Header().Add("content-type", "application/json")

	//get the database instance
	collection, ctx := data.GetDB(pr.l, "products")

	//Extract all records
	cursor, err := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	if err != nil {
		http.Error(rw, "Failed to get data from DB", http.StatusInternalServerError)
		return
	}

	//Collect all records for response
	var products []data.Product

	//Collect the records and store in slice
	for cursor.Next(ctx) {
		var product data.Product
		cursor.Decode(&product)
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		http.Error(rw, "Failed to get data from DB", http.StatusInternalServerError)
		return
	}
	//Write the response
	json.NewEncoder(rw).Encode(products)
}

// GetOneProduct This method calls one person in DB
func (pr *Products) GetOneProduct(rw http.ResponseWriter, r *http.Request) {
	pr.l.Println("\"************ GetOneProduct API ******************\"")

	//Update the response header for JSON
	rw.Header().Add("content-type", "application/json")

	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
		return
	}

	//get the database instance
	collection, ctx := data.GetDB(pr.l, "products")
	var product data.Product

	//Extract one record
	err = collection.FindOne(ctx, data.Product{ProdID: id}).Decode(&product)
	if err != nil {
		http.Error(rw, "Unable to parse data", http.StatusInternalServerError)
		return
	}

	//Write the response
	json.NewEncoder(rw).Encode(product)
}
