package tools

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBodyJson(r *http.Request, jsonReq interface{}) (interface{}, func() error, error) {
	err := json.NewDecoder(r.Body).Decode(jsonReq)
	if err != nil {
		return nil, nil, err
	}

	return jsonReq, r.Body.Close, nil
}
