package apis

import (
	"drone-task/usecase"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MedicationApis struct {
	MedicationUseCase usecase.IMedicationUseCase
}

func NewMedicationAPIS(MedicationUseCase usecase.IMedicationUseCase) MedicationApis {
	return MedicationApis{
		MedicationUseCase: MedicationUseCase,
	}
}

func (api MedicationApis) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error": "%s"}`, err.Error())))
		return
	}
	response, err := api.MedicationUseCase.Create(ctx, requestByte)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{ "error": %s}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
