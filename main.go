// soap_client project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
)

const (
	CONFIG_PATH string = `C:\goplace\src\config\soap_5\conf.ini`
)

type settings struct {
	Server server_settings  `toml:"server_settings"`
	Req    request_settings `toml:"request_settings"`
}

type server_settings struct {
	Access_data access_settings `toml:"access_settings"`
}

type access_settings struct {
	User   string
	Pwd    string
	Crypty string
	Time   int
}
type request_settings struct {
	Def_headers [][]string `toml:"default_headers"`
	Auth_xml    string     `toml:"auth_xml_pattern"`
	Auth_url    string     `toml:"auth_url"`
}

func (rs *request_settings) Def_headers_map() map[string]string {

	headers := map[string]string{}
	def_header_count := len(rs.Def_headers[0])
	for i := 0; i < def_header_count; i++ {
		headers[rs.Def_headers[0][i]] = rs.Def_headers[1][i]

	}
	return headers

}

var set settings

func main() {
	toml.DecodeFile(CONFIG_PATH, &set)
	post_req(set.Req.Auth_url, fmt.Sprintf(set.Req.Auth_xml, set.Server.Access_data.User, set.Server.Access_data.Pwd, set.Server.Access_data.Time, set.Server.Access_data.Crypty), set.Req.Def_headers_map())

}

func post_req(url string, body string, headers map[string]string) {
	client := &http.Client{}
	fmt.Printf("%s\r\n", body)
	req, _ := http.NewRequest("POST", url, strings.NewReader(body))
	for name, val := range headers {
		req.Header.Set(name, val)
	}
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(body)))
	resp, _ := client.Do(req)
	fmt.Println(resp.StatusCode)
	responseData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(responseData))
}
