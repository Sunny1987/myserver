package handlers

import (
	"context"
	"encoding/json"
	"myserver/myserver/data"
	"net/http"
)

type KeyPerson struct{}

// MiddleWareProductValidation This is the middleware support to add json checks and validatiions before handler updates
func (p *Persons) MiddleWareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			p.l.Println("************Middleware Validations start ******************")

			//Receive the input response from client and decode data
			var person data.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				http.Error(rw, "ERROR decoding the data", http.StatusBadRequest)
				return
			}

			//Provide the context for handler post middleware
			ctx := context.WithValue(r.Context(), KeyPerson{}, person)
			r = r.WithContext(ctx)

			next.ServeHTTP(rw, r)
			p.l.Println("************Middleware Validations ends ******************")
		},
	)
}
