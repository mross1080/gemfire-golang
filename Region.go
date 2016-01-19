package gemfireGolang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bytes"
)

type User struct {
	Name string
	id 	 string
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

func (api Api) getRegions(params map[string]string) ([]Region, int) {
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


func (api Api) getEntry(regionName string, key string) (map[string]string, int) {

	var entry map[string]string
	entry = make(map[string]string)
//	params := make(map[string]string)

	r, err := http.Get(api.Url() + regionName + "/" + key)
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

func (api Api) getRegion(regionName string, params map[string]string) (map[string][]map[string]string,int) {
	entry :=  make(map[string][]map[string]string)

	url := buildRequest(api.Url() + regionName,params)
	r, err := http.Get(url)
	fmt.Println("making get to ",api.Url() + regionName)
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


func (api Api) createEntry () {


	user := User{"Bob Ross","11"}
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	var jsonStr = []byte(string(u))

	resp, err := http.NewRequest("POST","http://127.0.0.1:8080/gemfire-api/v1/test/", bytes.NewBuffer(jsonStr))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}