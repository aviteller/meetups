package main

import (
	"fmt"
	"net/http"

	"../backend/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},                                               // All origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"}, // Allowing only get, just an example
	})

	router.HandleFunc("/api/meetups", controllers.CreateMeetup).Methods("POST")
	router.HandleFunc("/api/meetups", controllers.GetMeetups).Methods("GET")
	router.HandleFunc("/api/meetups/{id}", controllers.GetMeetup).Methods("GET")
	router.HandleFunc("/api/meetups/{id}", controllers.DeleteMeetup).Methods("DELETE")
	router.HandleFunc("/api/meetups/{id}", controllers.UpdateMeetup).Methods("PUT")
	router.HandleFunc("/api/meetups/togglelike/{id}", controllers.ToggleLike).Methods("PUT")

	// router.Use(app.JwtAuthentication)
	port := "9000"
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, c.Handler(router))
	if err != nil {
		fmt.Println(err)
	}

}
