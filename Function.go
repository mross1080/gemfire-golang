package gemfireGolang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

func (api Api) getFunctions() ([]string, int) {

//	params := make(map[string]string)

	r, err := http.Get(api.Url() + "functions")
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
