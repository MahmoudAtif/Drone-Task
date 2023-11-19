package main

import (
	"drone-task/http/apis"
	"drone-task/repository/db"
	"drone-task/repository/db/migrations"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := db.ConnectToDataBase()
	if err != nil {
		log.Println(err)
	}
	err = migrations.MigrateDrone(db)
	if err != nil {
		log.Println(err)
	}

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	dronesSubRouter := r.PathPrefix("/drones").Subrouter()
	myDroneAPI := apis.DroneAPIS{}
	dronesSubRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myDroneAPI.Get(w, r)
	}).Methods("GET")
	dronesSubRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myDroneAPI.Create(w, r)
	}).Methods("POST")

	http.ListenAndServe(":9999", r)

}
