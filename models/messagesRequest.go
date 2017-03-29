package models

type MessagesList struct {
	Messages []Message `json:"messages"`
}



type Message struct {
	DvcMsgId string `json:"dvcMsgId"`
	Name string `json:"name"`
	Text string `json:"text"`
	PhoneNo string `json:"phoneNo"`
	DateTime int64 `json:"dateTime"`
	MsgType string `json:"msgType"`
	AppType string `json:"appType"`
	ConvId string `json:"convId"`
	Operation string `json:"operation"`
}



type Invalid struct {
	Code        string `json:"code"`
	Error       string `json:"error"`
	DeviceMsgId string `json:"dvcMsgId"`
}

type Success struct {
	DeviceMsgId string `json:"dvcMsgId"`
	ServerId    string `json:"serMsgId"`
}