package types

import "errors"

//APIResponse holds the response from prometheus endpoint query
type APIResponse struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

//FilterAPIResponseValue receive the APIResponse and return the correspondent value
func (r *APIResponse) FilterAPIResponseValue() (string, error) {
	if len(r.Data["result"].([]interface{})) == 0 {
		return "", errors.New("No data received in the response")
	}

	return r.Data["result"].([]interface{})[0].(map[string]interface{})["value"].([]interface{})[1].(string), nil
}
