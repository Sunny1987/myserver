package handlers

import (
	"context"
	"encoding/json"
	"myserver/myserver/data"

	"net/http"
)

type KeyPerson struct{}

//type KeyProduct struct {}

// MiddleWareProductValidation This is the middleware support to add json checks and validations before handler updates
func (c *Persons) MiddleWareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			c.l.Println("************Middleware Validations start ******************")

			switch r.URL.Path {
			case "/add":
				//errPer,per := c.toJsonPerson(r)
				p := &data.Person{}
				err := json.NewDecoder(r.Body).Decode(p)
				c.l.Printf("The person is : %v", p)
				if err != nil {
					http.Error(rw, "ERROR decoding the data from MW for /add", http.StatusBadRequest)
					c.l.Printf("ERROR decoding the data : %v", err)
					return
				}
				ctx := context.WithValue(r.Context(), KeyPerson{}, p)
				r = r.WithContext(ctx)
				//next.ServeHTTP(rw, r)
			case "/update":
				p := &data.Person{}
				err := json.NewDecoder(r.Body).Decode(p)
				c.l.Printf("The person is : %v", p)
				if err != nil {
					http.Error(rw, "ERROR decoding the data from MW for /update", http.StatusBadRequest)
					c.l.Printf("ERROR decoding the data : %v", err)
					return
				}
				ctx := context.WithValue(r.Context(), KeyPerson{}, p)
				r = r.WithContext(ctx)
				//next.ServeHTTP(rw, r)
			case "/addprod":
				//errProd,prod := c.toJsonProduct(r)
				//err := c.toJsonPerson(r)
				pr := &data.Product{}
				err := json.NewDecoder(r.Body).Decode(pr)
				c.l.Printf("The product is : %v", pr)
				if err != nil {
					http.Error(rw, "ERROR decoding the data from MW for /addprod", http.StatusBadRequest)
					c.l.Printf("ERROR decoding the data : %v", err)
					return
				}
				ctx := context.WithValue(r.Context(), KeyPerson{}, pr)
				r = r.WithContext(ctx)
				//next.ServeHTTP(rw, r)
			case "/updateprod":
				pr := &data.Product{}
				err := json.NewDecoder(r.Body).Decode(pr)
				c.l.Printf("The product is : %v", pr)
				if err != nil {
					http.Error(rw, "ERROR decoding the data from MW for /updateprod", http.StatusBadRequest)
					c.l.Printf("ERROR decoding the data : %v", err)
					return
				}
				ctx := context.WithValue(r.Context(), KeyPerson{}, pr)
				r = r.WithContext(ctx)
			}
			c.l.Println("************Middleware Validations ends ******************")
			next.ServeHTTP(rw, r)

		},
	)
}
