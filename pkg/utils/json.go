package utils

import (
	"bytes"
	"encoding/json"
)

func PrettyJSON(data []byte) ([]byte, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "", "  ")

	if err != nil {
		return nil, err
	}

	return prettyJSON.Bytes(), nil
}
