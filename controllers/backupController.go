package controllers

import (
	"net/http"
	"fmt"
	"message_backup/models"
	"message_backup/validation"
	"message_backup/businessLogic"
	"encoding/json"
)


func MsgBackup(w http.ResponseWriter, r *http.Request) {

	deviceKey, userId, err := validation.ValidateHeaders(r)
	if err.Status != http.StatusOK {
		handleError(w, err)
		return
	}

	valid,invalid,partial,err := validation.RequestValidation(w, r)
	if err.Status != http.StatusOK {
		handleError(w, err)
		return
	}

	response, err := businessLogic.PutInCass(userId,deviceKey,valid)
	if err.Status != http.StatusOK {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if partial{
		response.Status = 2
		response.Invalid = invalid
		w.WriteHeader(http.StatusAccepted)

	} else {
		response.Status = 1
		response.Invalid = invalid
		w.WriteHeader(http.StatusOK)
	}
	res,errRes := json.Marshal(response)
	if errRes != nil {
		http.Error(w, errRes.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(res))

}

func handleError(w http.ResponseWriter, e models.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Status)
	response,err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, string(response))
	}
}