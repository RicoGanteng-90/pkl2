package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang-crud-clean-architecture/config/db"
	"golang-crud-clean-architecture/router"
	"log"
	"net/http"
)

func main() {
	db.InitialMigration()
	router := myrouter.NewRooter(mux.NewRouter())
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		AllowCredentials: true,
		Debug:            true,
	}).Handler(router)

	log.Fatal(http.ListenAndServe("localhost:8080", handler))

}

func AllowOriginFunc(r *http.Request, origin string) bool {
	if origin == "http://localhost:3000" {
		return true
	}
	return false
}
