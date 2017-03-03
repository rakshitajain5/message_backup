package models

type MessagesResponse struct {
	Status int `json:"status"`
	LastBackupTime int64 `json:"lastBackupTime"`
	LastMsgTime int64 `json:"lastMsgTime"`
	Success interface{} `json:"success"`
	Invalid []ErrorResponse `json:"invalid"`
}
