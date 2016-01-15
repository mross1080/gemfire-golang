package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"time"
)

type Api struct {
	hostName string
	port string
}

func (api Api) Url() string{
	return api.hostName + ":" + api.port + "/gemfire-api/v1/"

}

type Region struct {
	Name            string `json:"name"`
	RegionType      string `json:"type"`
	KeyConstraint   string `json:"key-constraint"`
	ValueConstraint string `json:"value-constraint"`
}

type ClusterRegions struct {
	Regions []Region
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


	api := Api{"http://127.0.0.1","8080"}
	fmt.Println(api.Url())
	rand.Seed(int64(time.Now().Nanosecond()))
	value := rand.Intn(2)
	fmt.Println(value)
	fmt.Println(api.getRegionKeys("designer"))


}

func getFunctions() ([]string, int) {

	r, err := http.Get("http://127.0.0.1:8080/gemfire-api/v1/functions")
	if err != nil {
		fmt.Println(err)
	} else {
		defer r.Body.Close()
		fmt.Println(r.Status)
		var m Functions
		requestBody, _ := ioutil.ReadAll(r.Body)
		e := json.Unmarshal(requestBody, &m)
		fmt.Println(e, m.Functions)
		fmt.Println("resoiobse is: ", r.StatusCode)
		return m.Functions, r.StatusCode
	}

	return nil, r.StatusCode

}


func iterateJson() {
	r, err := http.Get("http://127.0.0.1:8080/gemfire-api/v1")
	if err != nil {
		fmt.Println(err)
	} else {
		defer r.Body.Close()
		fmt.Println(r.Body)

		requestBody, _ := ioutil.ReadAll(r.Body)
		//		 var m ClusterRegions
		//		 e := json.Unmarshal(requestBody, &m)

		var f interface{}
		er := json.Unmarshal(requestBody, &f)
		marsh := f.(map[string]interface{})
		for k, v := range marsh {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					x := reflect.TypeOf(u)

					fmt.Println(i, x, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}

		fmt.Println(marsh, er)

	}
}

//func getEntry(regionName string, key string) (map[string]string, int) {
//
//	var entry map[string]string
//	entry = make(map[string]string)
//
//	r, err := http.Get("http://127.0.0.1:8080/gemfire-api/v1/" + regionName + "/" + key)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		defer r.Body.Close()
//		requestBody, _ := ioutil.ReadAll(r.Body)
//		e := json.Unmarshal(requestBody, &entry)
//		if e != nil {
//			fmt.Println(err)
//		}
//		fmt.Println("Returned", entry)
//
//		return entry, r.StatusCode
//	}
//
//	return entry, 200
//
//}
