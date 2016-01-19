package gemfireGolang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//	"net/url"
	"bytes"
	//	"reflect"
)

type User struct {
	Name string
	Id   string
}

type Region struct {
	api  Api
	Name string
}

func (api Api) getRegionKeys(regionName string) ([]string, int) {
	//	params := make(map[string]string)

	r, err := http.Get(api.Url() + regionName + "/keys")
	if err != nil {
		fmt.Println(err)
	} else {
		defer r.Body.Close()

		requestBody, _ := ioutil.ReadAll(r.Body)

		var a RegionKeys
		err = json.Unmarshal(requestBody, &a)
		return a.Keys, r.StatusCode

	}
	return nil, r.StatusCode
}

func (api Api) getRegions(params map[string]string) ([]RegionDef, int) {
	var m ClusterRegions

	r, err := http.Get(api.Url())
	if err != nil {
		fmt.Println(err)
	} else {
		defer r.Body.Close()

		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			fmt.Println(err)
		}

		return m.Regions, r.StatusCode
	}

	return m.Regions, r.StatusCode
}

func (region Region) get(key string) (map[string]string, int) {

	var entry map[string]string
	entry = make(map[string]string)
	//	params := make(map[string]string)

	r, err := http.Get(region.api.Url() + region.Name + "/" + key)
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

		return entry, r.StatusCode
	}

	return entry, 200

}

func (api Api) getRegion(regionName string, params map[string]string) (map[string][]map[string]string, int) {
	entry := make(map[string][]map[string]string)

	url := buildRequest(api.Url()+regionName, params)
	r, err := http.Get(url)
	fmt.Println("making get to ", api.Url()+regionName)
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

		return entry, r.StatusCode
	}

	return entry, 200

}

func (region Region) put(key string, js []uint8) int {

	url := region.api.Url() + region.Name + "/?key=" + key

	var body = []byte(string(js))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
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
