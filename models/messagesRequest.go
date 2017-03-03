package models

type MessagesList struct {
	Messages []Message `json:"messages"`
}

var Abc string = "sss";

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



