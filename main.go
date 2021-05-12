package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"myserver/myserver/data"
	"myserver/myserver/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	//create standard log
	l := log.New(os.Stdout, "server", log.LstdFlags)
	//personHandler := handlers.NewPerson(l)
	//productHandler := handlers.NewProduct(l)
	commonHandler := handlers.NewCommon(l)

	//create new server mux
	serverMux := mux.NewRouter()

	//Connect with DB
	data.DBConnect(l)

	//register handlers
	postRouter := serverMux.Methods("POST").Subrouter()
	getRouter := serverMux.Methods("GET").Subrouter()
	putRouter := serverMux.Methods("PUT").Subrouter()

	//*********************** Person **************************************//
	postRouter.HandleFunc("/add", commonHandler.Persons.CreatePOSTForPersons)
	getRouter.HandleFunc("/getall", commonHandler.Persons.GetAllPersons)
	getRouter.HandleFunc("/getone/{id}", commonHandler.Persons.GetOnePerson)
	putRouter.HandleFunc("/update/{id}", commonHandler.Persons.CreatePUTForPerson)

	//*********************** Person **************************************//

	//*********************** Product **************************************//
	postRouter.HandleFunc("/addprod", commonHandler.Products.CreatePOSTForProducts)
	getRouter.HandleFunc("/getallprod", commonHandler.Products.GetAllProducts)
	getRouter.HandleFunc("/getoneprod/{id}", commonHandler.Products.GetOneProduct)
	putRouter.HandleFunc("/updateprod/{id}", commonHandler.Products.CreatePUTForProduct)

	//*********************** Product **************************************//
	postRouter.Use(commonHandler.Persons.MiddleWareProductValidation)
	putRouter.Use(commonHandler.Persons.MiddleWareProductValidation)

	//create server
	myServer := &http.Server{
		Addr:         ":8080",
		Handler:      serverMux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     l,
	}

	//Set up the server to listen and Serve
	go func() {
		l.Println("Starting server at 8080")
		err := myServer.ListenAndServe()
		if err != nil {
			l.Printf("Error Starting the server: %s", err)
			os.Exit(1)
		}
	}()

	//set up the channel to read and notify user server kills/interrupts
	mySig := make(chan os.Signal, 1)
	signal.Notify(mySig, os.Interrupt)
	sig := <-mySig
	//sending the notification alert for server termination
	l.Println("The server has received a shutdown request.. Shutting down", sig)

	//Disconnect with DB
	//data.DBDisconnect(l)

	//server shutdown initiated within a time period of 30 seconds
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	myServer.Shutdown(tc)

}
