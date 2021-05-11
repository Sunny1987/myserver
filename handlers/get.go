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

	//set filter record criteria
	filter := bson.M{"_id": bson.M{"eq": id}}

	//update data

	result, err := collection.UpdateOne(ctx, filter, person)
	if err != nil {
		http.Error(rw, "Unable to update data", http.StatusBadRequest)
		p.l.Printf("Unable to update data: %v", err)
		return
	} else {
		p.l.Printf("Updated data: %v", result)
	}

	//Write the response
	json.NewEncoder(rw).Encode(result)
}
