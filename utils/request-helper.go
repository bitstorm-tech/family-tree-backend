package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ExtractFromRequest parses the request and tries to store the struct
// from the request body into the given type "to"
func ExtractFromRequest(request *http.Request, to interface{}) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, to)
	if err != nil {
		return err
	}

	return nil
}
