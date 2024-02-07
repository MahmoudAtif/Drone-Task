package apis

import (
	"drone-task/usecase"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DroneLoadAPIS struct {
	droneLoadUseCase usecase.IDroneLoadUseCase
}

func NewDroneLoadAPIS(droneLoadUseCase usecase.IDroneLoadUseCase) DroneLoadAPIS {
	return DroneLoadAPIS{
		droneLoadUseCase: droneLoadUseCase,
	}
}

func (api DroneLoadAPIS) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error": "%s"}`, err.Error())))
		return
	}
	response, err := api.droneLoadUseCase.Create(ctx, requestByte)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error": %s}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
