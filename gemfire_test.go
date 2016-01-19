package gemfireGolang

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"strconv"
//	"encoding/json"
//	"reflect"
)


func TestGetRegions(t *testing.T) {

	api := Api{"http://127.0.0.1","8080"}

	params := make(map[string]string)

	result, responseCode := api.getRegions(params)
	if result == nil {
		t.Fatalf("API response was nil")
	}

	if responseCode != 200 && responseCode != 300 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}


}

func TestGetRegion(t *testing.T) {

	api := Api{"http://127.0.0.1","8080"}
	regionName := "designer"
	params := make(map[string]string)
	entrylimit := 3

	params["limit"] = strconv.Itoa(entrylimit)

	entries, responseCode := api.getRegion(regionName,params)
	if len(entries[regionName]) > entrylimit {
		t.Fatalf("Set limit of 2 results got back ", len(entries))
	}

	if responseCode != 200 && responseCode != 300 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestGetFunctions(t *testing.T) {

	api := Api{"http://127.0.0.1","8080"}


	result, responseCode := api.getFunctions()
	if result == nil {
		t.Fatalf("API response was nil")
	}

	if responseCode != 200 && responseCode != 300 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestRegionKeys(t *testing.T) {

	api := Api{"http://127.0.0.1","8080"}

	params := make(map[string]string)
	regions, _ := api.getRegions(params)

	if len(regions) != 0 {
		rand.Seed(int64(time.Now().Nanosecond()))
		index := rand.Intn(2)
		regionName := regions[index].Name
		fmt.Println(regions)

		result, responseCode := api.getRegionKeys(regionName)
		if result == nil {
			t.Fatalf("API response was nil")
		}

		if responseCode != 200 && responseCode != 300 {
			t.Fatalf("Failed to hit api got response code of %i", responseCode)
		}

	}

}

func TestGetEntry(t *testing.T) {

	api := Api{"http://127.0.0.1", "8080"}

	params := make(map[string]string)





	regions, _ := api.getRegions(params)
	if len(regions) != 0 {
		rand.Seed(int64(time.Now().Nanosecond()))
		index := rand.Intn(len(regions))
		fmt.Println("Selected index %i for regions", index)
		regionName := regions[index].Name
		keys, _ := api.getRegionKeys(regionName)
		if len(keys) != 0 {
			rand.Seed(int64(time.Now().Nanosecond()))
			keyIndex := rand.Intn(len(keys))
			fmt.Println("Selected index %i for keys", keyIndex)

			key := keys[keyIndex]
			result, responseCode := api.getEntry(regionName, key)
			if result == nil {
				t.Fatalf("API response was nil")
			}

			if responseCode != 200 && responseCode != 300 {
				t.Fatalf("Failed to hit api got response code of %i", responseCode)
			}
		}


	}

}

func TestBuildRequest(t *testing.T) {
	baseUrl := "http://127.0.0.1:8080/gemfire-api/v1/designer/"
	params := make(map[string]string)
	params ["limit"] = "5"
	expected := "http://127.0.0.1:8080/gemfire-api/v1/designer/?limit=5"
	result := buildRequest(baseUrl,params)
	if result != expected {
		t.Fatalf("Did not build the URL to be %v got \n %v", expected,result)

	}

//	params["ignoreMissingKey"] = "true"
//	expected = "http://127.0.0.1:8080/gemfire-api/v1/designer/?limit=5&ignoreMissingKey=true"
//	result = buildRequest(baseUrl,params)
//	if result != expected {
//		t.Fatalf("Did not build the URL to be %v got \n %v", expected,result)
//
//	}



}

func TestCreateEntry(t *testing.T) {
//	baseUrl := "http://127.0.0.1:8080/gemfire-api/v1/test/"
//	params := make(map[string]string)
//
//	user := User{"Bob Ross","11"}
//	u, err := json.Marshal(user)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Print(reflect.TypeOf(u))
	api := Api{"http://127.0.0.1", "8080"}

//	params := make(map[string]string)
	api.createEntry()


}
