package bpm

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	//	"github.com/pborman/uuid"
)

type BPM struct {
	user    string
	pwd     string
	crypty  string
	time    int
	req_url string
	actions map[string]*Soap_actions
	token   string
	client  http.Client
}

func Init(user string, pwd string, crypty string, time int, req_url string, actions map[string]*Soap_actions) BPM {
	b := BPM{}
	b.user = user
	b.pwd = pwd
	b.crypty = crypty
	b.time = time
	b.req_url = req_url
	b.actions = actions
	b.client = http.Client{}
	b.get_token()
	return b
}

func (b *BPM) get_token() {
	act := "get_token"

	req_struct := get_login_struct()
	fmt.Println(req_struct.Body.Run_login.Xmlns)

	req := b.create_post_req(act, b.user, b.pwd, b.time, b.crypty)
	resp, _ := b.client.Do(req)
	responseData, _ := ioutil.ReadAll(resp.Body)
	get_token_struct := get_token{}
	xml.Unmarshal(responseData, &get_token_struct)
	if get_token_struct.Error_code == 0 {
		b.token = get_token_struct.Token
	}

}

func (b BPM) Select_data() {
	ra := envelope{}.init_select()
	fmt.Println(ra)
	by, _ := xml.Marshal(ra)
	fmt.Println(string(by))
}

func (b BPM) create_post_req(act string, i ...interface{}) *http.Request {
	req_body := fmt.Sprintf(b.actions[act].Body_ptr, i...)
	req, _ := http.NewRequest("POST", b.req_url, strings.NewReader(req_body))

	headers := b.actions[act].Def_headers_map()
	for name, val := range headers {
		req.Header.Set(name, val)
	}
	return req
}
