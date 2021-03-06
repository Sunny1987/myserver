package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myserver/myserver/data"
	"net/http"
)

// CreatePOSTForPerson This method is used to update existing record
func (p *Persons) CreatePOSTForPerson(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("\"************CreatePOSTForPerson API ******************\"")

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

	//get the person instance context post middleware
	person := r.Context().Value(KeyPerson{}).(data.Person)
	//var person data.Person

	//set filter to search data
	filter := bson.M{"_id": bson.M{"$eq": id}}

	//update := bson.M{person}
	update := bson.M{"$set": bson.M{
		"first_name": person.FirstName,
		"last_name":  person.LastName,
		"phone":      person.Phone,
		"email":      person.Email,
	}}

	// Call the driver's UpdateOne() method and pass filter and update to it
	result, err := collection.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		p.l.Printf("The update failed :%v", err)
		http.Error(rw, "Unable to update data", http.StatusInternalServerError)
	}

	//Write output response
	json.NewEncoder(rw).Encode(result)

}
