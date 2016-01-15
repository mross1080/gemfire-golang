package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetRegions(t *testing.T) {

	api := Api{"http://127.0.0.1","8080"}

	result, responseCode := api.getRegions()
	if result == nil {
		t.Fatalf("API response was nil")
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

	regions, _ := api.getRegions()

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

	api := Api{"http://127.0.0.1","8080"}

	regions, _ := api.getRegions()
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
