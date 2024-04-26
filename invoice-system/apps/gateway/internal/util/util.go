package util

import (
	"bytes"
	"encoding/json"
	"log"
)

func ConvertStructToBytes(input interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		log.Printf("Error marshalling input to JSON: %v", err)
		return nil, err
	}

	// Create request body
	//bodyBuffer := bytes.NewBuffer(jsonData)
	return jsonData, nil
}

func ConvertStructToBuffer(input interface{}) (*bytes.Buffer, error) {
	jsonData, err := ConvertStructToBytes(input)

	if err != nil {
		return nil, err
	}

	bodyBuffer := bytes.NewBuffer(jsonData)

	return bodyBuffer, nil
}
