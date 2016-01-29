package gemfireGolang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	"strings"
)

type User struct {
	Name string
	Id   string
}

type Region struct {
	Connection  Api
	Name string
}

func (region Region) GetKeys() ([]string, int) {

	r, err := http.Get(region.Connection.Url() + region.Name + "/keys")
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

func (connection Api) GetRegions() ([]RegionDef, int) {
	var m ClusterRegions

	r, err := http.Get(connection.Url())
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

func (region Region) Get(keys ...string) (map[string]interface{}, int) {

	var entry map[string]interface{}
	entry = make(map[string]interface{})
	//	params := make(map[string]string)

	url := region.Connection.Url() + region.Name + "/"
	for i, key := range keys {
		url += key
		if i != len(keys)-1 {
			url += ","
		}

	}
	fmt.Println(url)
	r, err := http.Get(url)
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

func (connection Api) GetRegion(regionName string) (map[string][]map[string]string, int) {
	entry := make(map[string][]map[string]string)

	url := connection.Url()+regionName
	r, err := http.Get(url)
	fmt.Println("making get to ", connection.Url()+regionName)
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

func (region Region) Put(key string, js []uint8) int {

	url := region.Connection.Url() + region.Name + "/?key=" + key

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

func (region Region) Clear() int {
	url := region.Connection.Url() + region.Name

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(b)
	fmt.Println("Cleared Region",region.Name)

	return resp.StatusCode
}

func (region Region) Delete(key string) int {
	url := region.Connection.Url() + region.Name + "/" + key

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(b))

	return resp.StatusCode
}

func (region Region) Update(key string, js []uint8) int {

	url := region.Connection.Url() + region.Name + "/" + key
	payload := strings.NewReader(string(js))
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return res.StatusCode
}
