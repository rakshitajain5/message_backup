package businessLogic

import (
	"time"
	"message_backup/models"
	"message_backup/resources"
	"crypto/sha1"
	"strconv"
	"github.com/gocql/gocql"
	//"message_backup/dal"
	"net/http"
)

var insert_messages_by_users string = "INSERT INTO messages_by_users (user_id,msg_hash,msg_time,address,app_type,category,conv_id,device_msg_id,last_updated_tx_stamp,msg_type,name,state,text) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
var update_messages_by_users string = "UPDATE messages_by_users SET created_tx_stamp = created_tx_stamp+?, device_key = device_key+? WHERE user_id =? AND msg_hash = ?"
var insert_activities_by_devices string = "INSERT INTO activities_by_devices (user_id,device_key,last_backup_time,last_msg_time) VALUES(?,?,?,?)"


func PutInCass(userId string, deviceKey string, msg []models.Message) (models.MessagesResponse, models.ErrorResponse){
	var maxMsgDateTime int64 = 0
	var lastBackUpTime int64 = 0
	var response models.MessagesResponse
	responseCodes := make([]map[string]interface{}, len(msg))
	batch := gocql.NewBatch(gocql.LoggedBatch)

	if(len(msg) == 0) {
		return response,models.ErrorResponse{resources.ERROR_CODE_ALL_INVALID_MESSAGES,"provide valid messages", http.StatusBadRequest}
	}
	for i:=0; i<len(msg); i++ {
		message := msg[i]
		//go func() {
		//	lastBackUpTime = time.Now().Unix()
		//	if message.DateTime > maxMsgDateTime {
		//		maxMsgDateTime = message.DateTime
		//	}
		//}()

		lastBackUpTime = time.Now().Unix()
		if message.DateTime > maxMsgDateTime {
			maxMsgDateTime = message.DateTime
		}
		var hash []byte
		hmac(message.Text, message.PhoneNo, message.DateTime, &hash)
		/*ToDo check for category*/
		batch.Query(insert_messages_by_users, userId, hash, message.DateTime, message.PhoneNo, message.AppType, "personal", message.ConvId, message.DvcMsgId, lastBackUpTime, message.MsgType, message.Name, message.Operation, message.Text)
		batch.Query(update_messages_by_users, []int64{lastBackUpTime}, []string{deviceKey}, userId, hash)
		responseCode := make(map[string]interface{})
		responseCode["dvcMsgId"] = message.DvcMsgId
		responseCode["serMsgId"] = hash
		responseCodes[i] = responseCode
	}
	batch.Query(insert_activities_by_devices, userId, deviceKey, lastBackUpTime, maxMsgDateTime)
	//err := dal.PushinCass(batch)
	//if err != nil{
	//	return response,models.ErrorResponse{resources.CASSANDRA_SERVER_ERROR, err.Error(), http.StatusInternalServerError}
	//}
	response.LastBackupTime = lastBackUpTime
	response.LastMsgTime = maxMsgDateTime
	response.Success = responseCodes
	return response,models.ErrorResponse{"","", http.StatusOK}
}

func hmac(text string, phoneNo string, msgTimestamp int64, hash *[]byte) {
	h := sha1.New()
	h.Write([]byte(resources.MESSAGE_HASH_KEY + text + phoneNo + strconv.FormatInt(msgTimestamp, 10)))
	*hash = h.Sum(nil)
}