package main

import (
	"drone-task/http/apis"
	"drone-task/repository"
	"drone-task/repository/db"
	"drone-task/repository/db/migrations"
	"drone-task/usecase"
	"fmt"
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
	err = migrations.MigrateMedication(db)
	if err != nil {
		log.Println(err)
	}
	err = migrations.MigrateDroneLoad(db)
	if err != nil {
		log.Println(err)
	}
	droneRepo := repository.NewDroneRepository(db)
	droneUseCase := usecase.NewDroneUseCase(droneRepo)
	droneApi := apis.NewDroneAPIS(droneUseCase)

	medicationRepo := repository.NewMedicationRepository(db)
	medicationUseCase := usecase.NewMedicationUseCase(medicationRepo)
	mediccationApi := apis.NewMedicationAPIS(medicationUseCase)

	droneLoadRepo := repository.NewDroneLoadRepository(db)
	droneLoadUseCase := usecase.NewDroneLoadUseCase(
		droneLoadRepo,
		droneRepo,
		medicationRepo,
	)
	droneLoadApi := apis.NewDroneLoadAPIS(droneLoadUseCase)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()

	dronesSubRouter := r.PathPrefix("/drone").Subrouter()
	dronesSubRouter.HandleFunc("/", droneApi.Get).Methods("GET")
	dronesSubRouter.HandleFunc("/", droneApi.Create).Methods("POST")
	dronesSubRouter.HandleFunc("/{id}/", droneApi.GetByid).Methods("GET")
	dronesSubRouter.HandleFunc("/{id}/", droneApi.Delete).Methods("DELETE")
	dronesSubRouter.HandleFunc("/", droneApi.Update).Methods("PUT")

	medicationSubRouter := r.PathPrefix("/medication").Subrouter()
	medicationSubRouter.HandleFunc("/", mediccationApi.Create).Methods("POST")

	droneLoadSubRouter := r.PathPrefix("/drone/load").Subrouter()
	droneLoadSubRouter.HandleFunc("/", droneLoadApi.Create).Methods("POST")

	fmt.Println("development server at http://127.0.0.1:9999")
	http.ListenAndServe(":9999", r)

}
