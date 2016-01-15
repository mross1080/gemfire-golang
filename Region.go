package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

func (api Api) getRegionKeys(regionName string) ([]string, int) {

	r, err := http.Get("http://127.0.0.1:8080/gemfire-api/v1/" + regionName + "/keys")
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

func (api Api) getRegions() ([]Region, int) {
	var m ClusterRegions

	r, err := http.Get("http://127.0.0.1:8080/gemfire-api/v1/")
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


func (api Api) getEntry(regionName string, key string) (map[string]string, int) {

	var entry map[string]string
	entry = make(map[string]string)

	r, err := http.Get("http://127.0.0.1:8080/gemfire-api/v1/" + regionName + "/" + key)
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
