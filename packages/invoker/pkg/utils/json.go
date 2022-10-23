package utils

import "encoding/json"

func ErrorJsonString(err error) string {
	body, err := json.Marshal(map[string]string{"error": err.Error()})
	if err != nil {
		return ""
	}

	return string(body)
}
