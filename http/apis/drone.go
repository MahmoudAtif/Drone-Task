package apis

import (
	"drone-task/usecase"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	log.Printf(fmt.Sprintf(`dfdf %v`, api.DroneUseCase))
	ctx := r.Context()
	response, err := api.DroneUseCase.Get(ctx)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error": %s}`, err.Error())))
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
		w.Write([]byte(fmt.Sprintf(`{ "error": "%s"}`, err.Error())))
		return
	}
	response, err := api.DroneUseCase.Create(ctx, requestByte)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{ "error": %s}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (api DroneAPIS) GetByid(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	strId := vars["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		err := errors.New("invalid ID")
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error": %q}`, err.Error())))
		return
	}
	response, err := api.DroneUseCase.GetById(ctx, id)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf(`{ "error": %s}`, "not found")))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (api DroneAPIS) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{error: %q}`, err.Error())))
		return
	}
	err = api.DroneUseCase.Delete(ctx, id)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{ "error": "%s"}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (api DroneAPIS) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error": "%s"}`, err.Error())))
	}
	response, err := api.DroneUseCase.Update(ctx, requestByte)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error": "%s"}`, err.Error())))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
