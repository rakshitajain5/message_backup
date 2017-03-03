package validation

import (
	"net/http"
	"message_backup/models"
	"encoding/json"
	"message_backup/resources"
	"fmt"
)

func ValidateHeaders(r *http.Request) (string,string,models.ErrorResponse) {
	deviceKey := r.Header.Get("X-Device-Key")
	userId := r.Header.Get("X-user-id")
	if deviceKey != "" && userId != ""{
		return deviceKey,userId,models.ErrorResponse{"","",http.StatusOK};
	}
	return deviceKey,userId,models.ErrorResponse{resources.ERROR_MSG_HEADER_MISSING, "device key or user-id header missing", http.StatusBadRequest}
}

func RequestValidation(w http.ResponseWriter, r *http.Request) ([]models.Message, []models.ErrorResponse, bool, models.ErrorResponse) {
	var msg models.MessagesList
	decodeJson := json.NewDecoder(r.Body)
	errRes := decodeJson.Decode(&msg)
	if errRes != nil{
		return nil,nil,false,models.ErrorResponse{resources.INVALID_JSON_OBJECT,"invalid json body", http.StatusBadRequest}
	}

	var err models.ErrorResponse
	err = jsonSignatureValidation(msg)
	if err.Status != http.StatusOK {
		return nil,nil,false,err
	}

	var invalid []models.ErrorResponse
	var valid []models.Message
	partial := false

	for i := 0; i < len(msg.Messages); i++ {
		checkMandatoryFieldErr := checkMandatoryFields(w, msg.Messages[i])
		if checkMandatoryFieldErr.Status == http.StatusOK{
			checkEnumValidationErr := checkEnumValidations(msg.Messages[i])
			if checkEnumValidationErr.Status == http.StatusOK {
				replaceValues(&msg.Messages[i])
				valid = append(valid, msg.Messages[i])
			} else {
				invalid = append(invalid, checkEnumValidationErr)
			}
		} else {
			invalid = append(invalid, checkMandatoryFieldErr)
		}
	}
	if len(valid) > 0 && len(invalid) > 0 {
		partial = true
	}
 	return valid,invalid,partial,models.ErrorResponse{"","", http.StatusOK}
}

func checkMandatoryFields(w http.ResponseWriter, msg models.Message) models.ErrorResponse{
	if msg.DvcMsgId == "" {
		return models.ErrorResponse{resources.ERROR_CODE_DEVICE_MSG_ID, "dvcMsgId not provided", http.StatusBadRequest}
	}
	if msg.PhoneNo == "" {
		return models.ErrorResponse{resources.ERROR_CODE_PHONE_NO, "Please provide valid phoneNo", http.StatusBadRequest}
	}
	if msg.Text == "" {
		return models.ErrorResponse{resources.ERROR_CODE_TEXT, "Please pass valid text field", http.StatusBadRequest}
	}
	if msg.DateTime == 0 {
		return models.ErrorResponse{resources.ERROR_CODE_DATE_TIME, "Please pass valid dateTime", http.StatusBadRequest}
	}
	if msg.MsgType == "" {
		return models.ErrorResponse{resources.ERROR_CODE_MSG_TYPE, "Please pass valid msgType", http.StatusBadRequest}
	}
	if msg.Operation == "" {
		return models.ErrorResponse{resources.ERROR_CODE_OPERATION, "Please pass valid operation", http.StatusBadRequest}
	}
	return models.ErrorResponse{"","", http.StatusOK}
}

func checkEnumValidations(msg models.Message) models.ErrorResponse{
	if msg.Operation == "add" || msg.Operation == "delete" {
		return models.ErrorResponse{"","", http.StatusOK}
	}
	return models.ErrorResponse{resources.ERROR_CODE_OPERATION,"operation value can only be add or delete", http.StatusBadRequest}
}

func replaceValues(msg *models.Message) {
	if msg.Operation == "add" {
		msg.Operation = "A"
	} else {
		msg.Operation = "D"
	}
}

func jsonSignatureValidation(msg models.MessagesList) models.ErrorResponse {
	if msg.Messages == nil {
		return models.ErrorResponse{resources.INVALID_JSON_OBJECT, "No messages were provided", http.StatusBadRequest}
	}
	if len(msg.Messages) == 0 {
		return models.ErrorResponse{resources.INVALID_JSON_OBJECT, "no message was provided in messages array", http.StatusBadRequest}
	}
	if len(msg.Messages) > resources.MESSAGE_BATCH_LIMIT {
		return models.ErrorResponse{resources.INVALID_JSON_OBJECT, "batch limit can not exceed", http.StatusBadRequest}
	}
	return models.ErrorResponse{"","", http.StatusOK}
}