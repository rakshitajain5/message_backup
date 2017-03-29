package validation

import (
	"net/http"
	"message_backup/models"
	"encoding/json"
	"message_backup/resources"
	"gopkg.in/gin-gonic/gin.v1"
)

func ValidateHeaders(c *gin.Context) (string,string,models.ErrorResponse) {
	deviceKey := c.Request.Header.Get("x-device-key")
	userId := c.Request.Header.Get("x-user-id")
	if deviceKey != "" && userId != ""{
		return deviceKey,userId,models.ErrorResponse{"","",http.StatusOK};
	}
	return deviceKey,userId,models.ErrorResponse{"", "", http.StatusBadRequest}
}


func JsonStructureValidation(c *gin.Context)(models.MessagesList,models.ErrorResponse){
	var msg models.MessagesList
	decodeJson := json.NewDecoder(c.Request.Body)
	errRes := decodeJson.Decode(&msg)
	if errRes != nil{
		return msg,models.ErrorResponse{"","", http.StatusBadRequest}
	}
	return msg,models.ErrorResponse{"","", http.StatusOK}
}


func JsonSignatureValidation(msg models.MessagesList) (models.MessagesList,models.ErrorResponse) {
	if msg.Messages == nil {
		return msg,models.ErrorResponse{"", "No messages were provided", http.StatusBadRequest}
	}
	if len(msg.Messages) == 0 {
		return msg,models.ErrorResponse{"", "no message was provided in messages array", http.StatusBadRequest}
	}
	if len(msg.Messages) > resources.MESSAGE_BATCH_LIMIT {
		return msg,models.ErrorResponse{"", "batch limit can not exceed", http.StatusBadRequest}
	}
	return msg,models.ErrorResponse{"","", http.StatusOK}
}


func RequestValidation(msg models.MessagesList) ([]models.Message, []models.Invalid, bool, models.ErrorResponse) {
        length:= len(msg.Messages)
	var invalid []models.Invalid
	var valid []models.Message
	partial := false

	for i := 0; i < length; i++ {
		checkMandatoryFieldErr := checkMandatoryFields(msg.Messages[i])
		if checkMandatoryFieldErr.Status == http.StatusOK{
			checkEnumValidationErr := checkEnumValidations(msg.Messages[i])
			if checkEnumValidationErr.Status == http.StatusOK {
				replaceValues(&msg.Messages[i])
				valid = append(valid, msg.Messages[i])
			} else {
				invalid = append(invalid,models.Invalid{checkEnumValidationErr.Code,checkEnumValidationErr.Error,msg.Messages[i].DvcMsgId})
			}
		} else {
			invalid = append(invalid,models.Invalid{checkMandatoryFieldErr.Code,checkMandatoryFieldErr.Error,msg.Messages[i].DvcMsgId})
		}
	}
	if len(valid) > 0 && len(invalid) > 0 {
		partial = true
	}
 	return valid,invalid,partial,models.ErrorResponse{"","", http.StatusOK}
}

func checkMandatoryFields(msg models.Message) models.ErrorResponse{
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

