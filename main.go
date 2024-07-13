package main

import (
	"log"
	"net/http"
	"time"

	"github.com/granadosbrand/go-courses-backend/internal/user"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// port := ":3001"
	userEnd := user.MakeEndpoints()
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")
	
	// router.HandleFunc("/courses", getCourses)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
