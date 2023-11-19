package apis

import (
	"drone-task/usecase"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DroneAPIS struct {
	DroneUseCase usecase.IDroneUseCase
}

func NewDroneAPIS(DroneUseCase usecase.IDroneUseCase) DroneAPIS {
	return DroneAPIS{
		DroneUseCase: DroneUseCase,
	}
}

func (api DroneAPIS) Get(w http.ResponseWriter, r *http.Request) {
	log.Printf("Apiiiiiiiiiiiiiiiiiiiiiiiii")
	ctx := r.Context()
	response, err := api.DroneUseCase.Get(ctx)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error" : %s}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (api DroneAPIS) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}
	response, err := api.DroneUseCase.Create(ctx, requestByte)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{ "error" : %s}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
