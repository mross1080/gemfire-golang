package gemfireGolang

import (
	"fmt"
	"math/rand"
	"time"
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
