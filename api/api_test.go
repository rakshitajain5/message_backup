package api_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"message_backup/api"
	"strings"
	"net/http"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader //Ignore this for now
	usersUrl string
)

func init() {
	//fmt.Println(Handlers().GetRoute("/hello"))
	//server = httptest.NewServer(api.) //Creating new server with the user handlers
	server = httptest.NewServer(api.Handlers())
	usersUrl = fmt.Sprintf("%s/jcm/messages/backup", server.URL) //Grab the address for the API endpoint
}


func TestMessageBackup(t *testing.T) {
	//userJson := `{"username": "dennis", "balance": 200}`
	//p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	msgJson := `{"messages":[{"dvcMsgId":"4", "name":"jhsgvj", "text":"abc", "phoneNo":"9920183969", "dateTime":1257894000, "msgType":"1", "appType":"e", "convId":"232", "operation":"delete"}]}`

	reader = strings.NewReader(msgJson) //Convert string to reader

	request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-user-id", "test-userid-1")
	request.Header.Set("X-Device-Key", "test-devicekey-2")

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}

	//p.Stop()
}