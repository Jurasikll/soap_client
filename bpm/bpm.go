package bpm

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	//	"strings"
	"bytes"

	//	"github.com/pborman/uuid"
)

type BPM struct {
	user    string
	pwd     string
	crypty  string
	soap    string
	xmlns   string
	time    int
	req_url string
	actions map[string]*Soap_actions
	token   string
	client  http.Client
	ra      env_entity
}

func Init(user string, pwd string, crypty string, time int, req_url string, soap string, xmlns string, actions map[string]*Soap_actions) BPM {
	b := BPM{}
	b.user = user
	b.pwd = pwd
	b.crypty = crypty
	b.time = time
	b.req_url = req_url
	b.actions = actions
	b.soap = soap
	b.xmlns = xmlns
	b.ra = env_entity{}
	b.client = http.Client{}
	b.get_token()
	return b
}

func (b *BPM) get_token() {
	act := "get_token"
	b.ra.create_login(b.soap, b.crypty, b.time, b.user, b.pwd, b.xmlns)
	req := b.create_post_req(act, b.ra.get_xml())
	resp, _ := b.client.Do(req)
	responseData, _ := ioutil.ReadAll(resp.Body)
	get_token_struct := get_token{}
	xml.Unmarshal(responseData, &get_token_struct)
	if get_token_struct.Error_code == 0 {
		b.token = get_token_struct.Token
	}

}

func (b BPM) Select_data() {
	ra := env_entity{}
	ra.create_login(b.soap, b.crypty, b.time, b.user, b.pwd, b.xmlns)
	fmt.Printf("%v\r\n", ra)
	fmt.Println(ra.get_xml())
}

func (b BPM) create_post_req(act string, body []byte) *http.Request {
	//	req_body := fmt.Sprintf(b.actions[act].Body_ptr, i...)
	req, _ := http.NewRequest("POST", b.req_url, bytes.NewReader(body))

	headers := b.actions[act].Def_headers_map()
	for name, val := range headers {
		req.Header.Set(name, val)
	}
	return req
}
