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
	fmt.Println(api.getFunctions())


}


