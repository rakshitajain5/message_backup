package models

type MessagesResponse struct {
	Status int8 `json:"status"`
	Errors []string `json:"errors"`
	LastBackupTime int64 `json:"lastBackupTime"`
	LastMsgTime int64 `json:"lastMsgTime"`
	Success []Success `json:"success"`
	Invalid []Invalid `json:"invalid"`
}
