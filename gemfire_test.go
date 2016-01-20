package gemfireGolang

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	//	"reflect"
)

func TestGetRegions(t *testing.T) {

	api := Api{"http://127.0.0.1", "8080"}

	params := make(map[string]string)

	result, responseCode := api.GetRegions(params)
	if result == nil {
		t.Fatalf("API response was nil")
	}
	fmt.Println(result)
	if responseCode != 200 && responseCode != 300 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestGetRegion(t *testing.T) {

	api := Api{"http://127.0.0.1", "8080"}
	regionName := "designer"
	params := make(map[string]string)
	entrylimit := 3

	params["limit"] = strconv.Itoa(entrylimit)

	entries, responseCode := api.GetRegion(regionName, params)
	if len(entries[regionName]) > entrylimit {
		t.Fatalf("Set limit of 2 results got back ", len(entries))
	}

	if responseCode != 200 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestGetFunctions(t *testing.T) {

	api := Api{"http://127.0.0.1", "8080"}

	result, responseCode := api.getFunctions()
	if result == nil {
		t.Fatalf("API response was nil")
	}

	if responseCode != 200 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestGetKeysForRegion(t *testing.T) {

	api := Api{"http://127.0.0.1", "8080"}
	region := Region{api, "test"}

	result, responseCode := api.GetRegionKeys(region.Name)
	if result == nil {
		t.Fatalf("API response was nil")
	}

	fmt.Println(result)

	if responseCode != 200 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestGetEntry(t *testing.T) {

	api := Api{"http://127.0.0.1", "8080"}

	//	params := make(map[string]string)

	region := Region{api, "test"}
	region.Clear()

	user := User{"Freddy", "ad"}
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	responseCode := region.Put(user.Id, u)

	if responseCode != 200 && responseCode != 201 {
		t.Fatalf("Failed to hit api got response code of %v", responseCode)
	}

	result, responseCode := region.Get(user.Id)
	if result == nil {
		t.Fatalf("API response was nil")
	}

	if responseCode != 200 && responseCode != 300 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}
func TestGetEntries(t *testing.T) {

	api := Api{"http://127.0.0.1", "8080"}

	//	params := make(map[string]string)

	region := Region{api, "test"}
	region.Clear()

	user := User{"Freddy", "34"}
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	responseCode := region.Put(user.Id, u)


	user2 := User{"Sally", "33"}
	u2, err := json.Marshal(user2)
	if err != nil {
		fmt.Println(err)
	}
	responseCode = region.Put(user2.Id, u2)

	result, responseCode := region.Get(user.Id,user2.Id)
	if result == nil {
		t.Fatalf("API response was nil")
	}

	if responseCode != 200 && responseCode != 300 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}


}

func TestBuildRequest(t *testing.T) {
	baseUrl := "http://127.0.0.1:8080/gemfire-api/v1/designer/"
	params := make(map[string]string)
	params["limit"] = "5"
	expected := "http://127.0.0.1:8080/gemfire-api/v1/designer/?limit=5"
	result := buildRequest(baseUrl, params)
	if result != expected {
		t.Fatalf("Did not build the URL to be %v got \n %v", expected, result)

	}

}

func TestCreateEntry(t *testing.T) {

	api := Api{"http://127.0.0.1", "8080"}
	region := Region{api, "test"}

	user := User{"Freddy", "1asd"}
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	region.Put(user.Id, u)

}

func TestDeleteEntries(t *testing.T) {
	api := Api{"http://127.0.0.1", "8080"}
	region := Region{api, "test"}

	responseCode := region.Clear()

	if responseCode != 200 && responseCode != 201 {
		t.Fatalf("Failed to hit api got response code of %v", responseCode)
	}

}

func TestDeleteEntry(t *testing.T) {
	api := Api{"http://127.0.0.1", "8080"}
	region := Region{api, "test"}

	user := User{"Freddy", "1asd"}
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	region.Put(user.Id, u)

	responseCode := region.Delete(user.Id)

	if responseCode != 200 && responseCode != 201 {
		t.Fatalf("Failed to hit api got response code of %v", responseCode)
	}
}

func TestUpdateEntry(t *testing.T) {
	api := Api{"http://127.0.0.1", "8080"}
	region := Region{api, "test"}

	user := User{"Freddy", "aa"}
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	region.Put(user.Id, u)
	user.Name = "Taye Diggs"

	u, err = json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	responseCode := region.Update(user.Id, u)

	if responseCode != 200 && responseCode != 201 {
		t.Fatalf("Failed to hit api got response code of %v", responseCode)
	}

}
