package validation

import (
	"net/http"
	"message_backup/models"
	"encoding/json"
	"message_backup/resources"
	//"sync"
)

func ValidateHeaders(r *http.Request, deviceKey *string, userId *string, c chan models.ErrorResponse) {
	*deviceKey = r.Header.Get("X-Device-Key")
	*userId = r.Header.Get("X-user-id")
	if *deviceKey != "" && *userId != ""{
		//*err = models.ErrorResponse{"","",http.StatusOK}
		c <- models.ErrorResponse{"", "", http.StatusOK}
	} else {
		//*err = models.ErrorResponse{resources.ERROR_MSG_HEADER_MISSING, "device key or user-id header missing", http.StatusBadRequest}
		c <- models.ErrorResponse{resources.ERROR_MSG_HEADER_MISSING, "device key or user-id header missing", http.StatusBadRequest}
	}
	//wg.Done()
}

func RequestValidation(r *http.Request, valid *[]models.Message, invalid *[]models.ErrorResponse, partial *bool, c chan models.ErrorResponse) {
	var msg models.MessagesList
	decodeJson := json.NewDecoder(r.Body)
	errRes := decodeJson.Decode(&msg)
	var err1 models.ErrorResponse
	var err2 models.ErrorResponse
	var err3 models.ErrorResponse
	if errRes != nil{
		//*err = models.ErrorResponse{resources.INVALID_JSON_OBJECT,"invalid json body", http.StatusBadRequest}
		c <- models.ErrorResponse{resources.INVALID_JSON_OBJECT,"invalid json body", http.StatusBadRequest}
		return
	}

	//var wg1 sync.WaitGroup
	//wg1.Add(1)
	c1 := make(chan models.ErrorResponse)
	go jsonSignatureValidation(msg, c1)
	c2 := make(chan models.ErrorResponse)
	c3 := make(chan models.ErrorResponse)

	for i := 0; i < len(msg.Messages); i++ {
		//var wg2 sync.WaitGroup
		//wg2.Add(2)
		go checkMandatoryFields(msg.Messages[i], c2)
		go checkEnumValidations(msg.Messages[i], c3)
		//wg2.Wait()
		err2, err3 = <-c2, <-c3
		if err2.Status == http.StatusOK && err3.Status == http.StatusOK {
			*valid = append(*valid, msg.Messages[i])
		} else {
			if err2.Status != http.StatusOK {
				*invalid = append(*invalid, err2)
			} else {
				*invalid = append(*invalid, err3)
			}
		}
	}
	if len(*valid) > 0 && len(*invalid) > 0 {
		*partial = true
	}

	//wg1.Wait()
	err1 = <-c1
	if err1.Status != http.StatusOK {
		//*err = err1
		c <- err1
	} else {
		//*err = models.ErrorResponse{"","", http.StatusOK}
		c <- models.ErrorResponse{"","", http.StatusOK}
	}
	close(c1)
	close(c2)
	close(c3)
	//wg.Done()
}

func checkMandatoryFields(msg models.Message, c chan models.ErrorResponse) {
	if msg.DvcMsgId == "" {
		//*err = models.ErrorResponse{resources.ERROR_CODE_DEVICE_MSG_ID, "dvcMsgId not provided", http.StatusBadRequest}
		//wg.Done()
		c <- models.ErrorResponse{resources.ERROR_CODE_DEVICE_MSG_ID, "dvcMsgId not provided", http.StatusBadRequest}
		return
	}
	if msg.PhoneNo == "" {
		//*err = models.ErrorResponse{resources.ERROR_CODE_PHONE_NO, "Please provide valid phoneNo", http.StatusBadRequest}
		//wg.Done()
		c <- models.ErrorResponse{resources.ERROR_CODE_PHONE_NO, "Please provide valid phoneNo", http.StatusBadRequest}
		return
	}
	if msg.Text == "" {
		//*err = models.ErrorResponse{resources.ERROR_CODE_TEXT, "Please pass valid text field", http.StatusBadRequest}
		//wg.Done()
		c <- models.ErrorResponse{resources.ERROR_CODE_TEXT, "Please pass valid text field", http.StatusBadRequest}
		return
	}
	if msg.DateTime == 0 {
		//*err = models.ErrorResponse{resources.ERROR_CODE_DATE_TIME, "Please pass valid dateTime", http.StatusBadRequest}
		//wg.Done()
		c <- models.ErrorResponse{resources.ERROR_CODE_DATE_TIME, "Please pass valid dateTime", http.StatusBadRequest}
		return
	}
	if msg.MsgType == "" {
		//*err = models.ErrorResponse{resources.ERROR_CODE_MSG_TYPE, "Please pass valid msgType", http.StatusBadRequest}
		//wg.Done()
		c <- models.ErrorResponse{resources.ERROR_CODE_MSG_TYPE, "Please pass valid msgType", http.StatusBadRequest}
		return
	}
	if msg.Operation == "" {
		//*err = models.ErrorResponse{resources.ERROR_CODE_OPERATION, "Please pass valid operation", http.StatusBadRequest}
		//wg.Done()
		c <- models.ErrorResponse{resources.ERROR_CODE_OPERATION, "Please pass valid operation", http.StatusBadRequest}
		return
	}
	//*err = models.ErrorResponse{"","", http.StatusOK}
	//wg.Done()
	c <- models.ErrorResponse{"","", http.StatusOK}
}

func checkEnumValidations(msg models.Message, c chan models.ErrorResponse) {
	if msg.Operation == "add" || msg.Operation == "delete" {
		replaceValues(&msg)
		//*err = models.ErrorResponse{"","", http.StatusOK}
		c <- models.ErrorResponse{"","", http.StatusOK}
	} else {
		//*err = models.ErrorResponse{resources.ERROR_CODE_OPERATION,"operation value can only be add or delete", http.StatusBadRequest}
		c <- models.ErrorResponse{resources.ERROR_CODE_OPERATION,"operation value can only be add or delete", http.StatusBadRequest}
	}
	//wg.Done()
}

func replaceValues(msg *models.Message) {
	if msg.Operation == "add" {
		msg.Operation = "A"
	} else {
		msg.Operation = "D"
	}
}

func jsonSignatureValidation(msg models.MessagesList, c chan models.ErrorResponse) {
	if msg.Messages == nil {
		//*err = models.ErrorResponse{resources.INVALID_JSON_OBJECT, "No messages were provided", http.StatusBadRequest}
		c <- models.ErrorResponse{resources.INVALID_JSON_OBJECT, "No messages were provided", http.StatusBadRequest}
		//wg.Done()
		return
	}
	if len(msg.Messages) == 0 {
		//*err = models.ErrorResponse{resources.INVALID_JSON_OBJECT, "no message was provided in messages array", http.StatusBadRequest}
		c <- models.ErrorResponse{resources.INVALID_JSON_OBJECT, "no message was provided in messages array", http.StatusBadRequest}
		//wg.Done()
		return
	}
	if len(msg.Messages) > resources.MESSAGE_BATCH_LIMIT {
		//*err = models.ErrorResponse{resources.INVALID_JSON_OBJECT, "batch limit can not exceed", http.StatusBadRequest}
		c <- models.ErrorResponse{resources.INVALID_JSON_OBJECT, "batch limit can not exceed", http.StatusBadRequest}
		//wg.Done()
		return
	}
	//*err = models.ErrorResponse{"","", http.StatusOK}
	c <- models.ErrorResponse{"","", http.StatusOK}
	//wg.Done()
}