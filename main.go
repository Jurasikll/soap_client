// soap_client project main.go
package main

import (
	"fmt"

	"soap_client/bpm"

	"github.com/BurntSushi/toml"
)

const (
	CONFIG_PATH string = `C:\goplace\src\config\soap_5\conf.ini`
)

type settings struct {
	Conn_settings conn_settings
	Req           request_settings             `toml:"request_settings"`
	Soap_act      map[string]*bpm.Soap_actions `toml:"soap_actions"`
}

type conn_settings struct {
	User    string
	Pwd     string
	Crypty  string
	Time    int
	Req_url string
}

type request_settings struct {
	Def_headers   [][]string `toml:"default_headers"`
	Auth_xml      string     `toml:"auth_xml_pattern"`
	Auth_url      string     `toml:"auth_url"`
	Base_auth_url string
}

var set settings

func main() {
	toml.DecodeFile(CONFIG_PATH, &set)
	fmt.Scanln()
	bpm_client := bpm.Init(set.Conn_settings.User, set.Conn_settings.Pwd, set.Conn_settings.Crypty, set.Conn_settings.Time, set.Conn_settings.Req_url, set.Soap_act)
	bpm_client.Select_data()
	//	post_req(set.Req.Auth_url, fmt.Sprintf(set.Req.Auth_xml, set.Server.Access_data.User, set.Server.Access_data.Pwd, set.Server.Access_data.Time), set.Req.Def_headers_map())
	//	test_simple_req()
	fmt.Scanln()
	//test_base_auth()
}

func test_simple_req() {
	//	req_body := fmt.Sprintf(set.Req.Auth_xml, set.Server.Access_data.User, set.Server.Access_data.Pwd, set.Server.Access_data.Time, set.Server.Access_data.Crypty)
	//	req, _ := http.NewRequest("POST", set.Req.Auth_url, strings.NewReader(req_body))

	//	//	headers := set.Req.Def_headers_map()

	//	//	for name, val := range headers {
	//	//		req.Header.Set(name, val)
	//	//	}

	//	client := &http.Client{}
	//	resp, _ := client.Do(req)
	//	fmt.Println(resp.StatusCode)
	//	responseData, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Println(string(responseData))

}
