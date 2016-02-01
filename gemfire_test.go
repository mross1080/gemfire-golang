package gemfireGolang

import (
	"encoding/json"
	"fmt"
	"testing"
	//	"reflect"
)

func TestGetRegions(t *testing.T) {

	api := Api{"http://127.0.0.1", "8081"}

	result, responseCode := api.GetRegions()
	if result == nil {
		t.Fatalf("API response was nil")
	}
	fmt.Println(result)
	if responseCode != 200 && responseCode != 300 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestGetRegion(t *testing.T) {

	api := Api{"http://127.0.0.1", "8081"}
	regionName := "test"
	entries, responseCode := api.GetRegion(regionName)

	if responseCode != 200 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

	fmt.Println(entries)

}

func TestGetFunctions(t *testing.T) {

	api := Api{"http://127.0.0.1", "8081"}

	result, responseCode := api.getFunctions()
	if result == nil {
		t.Fatalf("API response was nil")
	}

	if responseCode != 200 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestGetKeysForRegion(t *testing.T) {

	api := Api{"http://127.0.0.1", "8081"}
	region := Region{api, "test"}

	result, responseCode := region.GetKeys()
	if result == nil {
		t.Fatalf("API response was nil")
	}

	fmt.Println(result)

	if responseCode != 200 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestGetEntry(t *testing.T) {

	api := Api{"http://127.0.0.1", "8081"}

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

	api := Api{"http://127.0.0.1", "8081"}

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

	result, responseCode := region.Get(user.Id, user2.Id)
	if result == nil {
		t.Fatalf("API response was nil")
	}

	if responseCode != 200 && responseCode != 300 {
		t.Fatalf("Failed to hit api got response code of %i", responseCode)
	}

}

func TestBuildRequest(t *testing.T) {
	baseUrl := "http://127.0.0.1:8081/gemfire-api/v1/designer/"
	params := make(map[string]string)
	params["limit"] = "5"
	expected := "http://127.0.0.1:8081/gemfire-api/v1/designer/?limit=5"
	result := buildRequest(baseUrl, params)
	if result != expected {
		t.Fatalf("Did not build the URL to be %v got \n %v", expected, result)

	}

}

func TestCreateEntry(t *testing.T) {

	api := Api{"http://127.0.0.1", "8081"}
	region := Region{api, "test"}

	region.Clear()
	user := struct {
		Name    string
		Age     string
		Id      string
		Company string
	}{
		"Bugsy Siegel",
		"22",
		"100",
		"Carrot Company"}

	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	region.Put(user.Id, u)
	entry, _ := region.Get(user.Id)
	if len(entry) == 0 {
		t.Fatalf("Failed to put the object in the region")
	}
	fmt.Println("result is ", entry)

}

func TestDeleteEntries(t *testing.T) {
	api := Api{"http://127.0.0.1", "8081"}
	region := Region{api, "test"}

	//	responseCode := region.Clear()

	regionEntires, responseCode := api.GetRegion(region.Name)

	fmt.Println(regionEntires)

	if responseCode != 200 && responseCode != 201 {
		t.Fatalf("Failed to hit api got response code of %v", responseCode)
	}

}

func TestDeleteEntry(t *testing.T) {
	api := Api{"http://127.0.0.1", "8081"}
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
	api := Api{"http://127.0.0.1", "8081"}
	region := Region{api, "test"}

	//	user := User{"Freddy", "aa"
	user := struct {
		Name string
		Age  string
		Id   string
	}{
		"Bobby Booshay",
		"234",
		"12"}
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	region.Put(user.Id, u)
	user.Name = "Taye Diggddds"

	u, err = json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	responseCode := region.Update(user.Id, u)

	if responseCode != 200 && responseCode != 201 {
		t.Fatalf("Failed to hit api got response code of %v", responseCode)
	}

}

func TestAdhocQuery(t *testing.T) {
	api := Api{"http://127.0.0.1", "8081"}
	region := Region{api, "test"}

	user := struct {
		Name string
		Age  string
		Id   string
	}{
		"Bobby Booshay",
		"234",
		"1552"}
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	region.Put(user.Id, u)
	result, _ := region.Get(user.Id)
	if result == nil {
		t.Fatalf("Failed to seed database")
	}

	queryString := "select Name,Age from /test where Id='1552'"

	queryResults, responseCode := api.AdHocQuery(queryString)

	if responseCode != 200 {
		t.Fatalf("Failed to execute query")

	}

	fmt.Println(queryResults)

}

func TestRegisterQuery(t *testing.T) {
	api := Api{"http://127.0.0.1", "8081"}

	queryString := "select Name,Id from /test"
	responseCode := api.RegisterQuery("1223", queryString)

	if responseCode != 201 {
		t.Fatalf("Failed to execute query")

	}

}

func TestExecuteQuery(t *testing.T) {

	api := Api{"http://127.0.0.1", "8081"}

	_, responseCode := api.ExecuteQuery("123", "")

	if responseCode != 200 {
		t.Fatalf("Failed to execute query")

	}

	fmt.Println("asdf")

}
