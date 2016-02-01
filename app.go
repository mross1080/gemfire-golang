package gemfireGolang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
	"bytes"
)

type Api struct {
	HostName string
	Port     string
}

func (api Api) Url() string {
	return api.HostName + ":" + api.Port + "/gemfire-api/v1/"
}

type RegionDef struct {
	Name            string `json:"name"`
	RegionType      string `json:"type"`
	KeyConstraint   string `json:"key-constraint"`
	ValueConstraint string `json:"value-constraint"`
}

type ClusterRegions struct {
	Regions []RegionDef
}

type RegionKeys struct {
	Keys []string
}

type Functions struct {
	Functions []string
}

func main() {
	//	getRegions()
	//	getFunctions()
	//		getRegionKeys("designer")2

	//	iterateJson()

	//	params := make(map[string]string)

	api := Api{"http://127.0.0.1", "8080"}
	fmt.Println(api.Url())
	rand.Seed(int64(time.Now().Nanosecond()))
	value := rand.Intn(2)
	fmt.Println(value)
	fmt.Println(api.getFunctions())

}

func buildRequest(baseUrl string, params map[string]string) string {

	count := 1
	if len(params) > 0 {
		baseUrl += "?"
		for key, param := range params {
			baseUrl += key + "=" + param
			if count != len(params) {
				baseUrl += "&"
				count += 1
			}

		}
		fmt.Println(baseUrl)
	}
	return baseUrl
}

func (connection Api) AdHocQuery(queryString string) ([]interface{}, int) {

	qs, _ := Encode(queryString)
	url := connection.Url() + "queries/adhoc?q=" + qs
	r, err := http.Get(url)
	var entry []interface{}
	fmt.Println("making get to ", url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer r.Body.Close()
		requestBody, _ := ioutil.ReadAll(r.Body)
		e := json.Unmarshal(requestBody, &entry)
		if e != nil {
			fmt.Println(err)
		}
		fmt.Println("Returned", entry)

		return entry, 200
	}

	return entry, 200

}

func (connection Api) RegisterQuery(id string, queryString string) int {

	qs, _ := Encode(queryString)

	url := connection.Url() + "queries/?id=" + id + "&q=" + qs

	fmt.Println("Registering query to ",url)
	fmt.Println("queryString is",url)
	var body = []byte(string(""))
	req, err := http.NewRequest("POST", url,  bytes.NewBuffer(body))
	req.Header.Set("X-Custom-Header", "entry-value")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(b))

	return resp.StatusCode
}

func (connection Api) ExecuteQuery(id string, params string) ([]interface{}, int) {

	if params == "" {
		params = "[\n    {\"@type\":\"int\",\n        \"@value\":2\n    }\n    ]"
	}

	url := connection.Url() + "queries/" + id
	payload := strings.NewReader(string(params))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	var results []interface{}

	if err != nil {
		panic(err)
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		e := json.Unmarshal(body, &results)
		if e != nil {
			fmt.Println(err)
		}
		fmt.Println("Returned", results)

		return results, 200

	}

	var a []interface{}

	return a, res.StatusCode

}

func Encode(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
