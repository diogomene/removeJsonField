package util

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadJsonFile(filePath string) (interface{}, error) {

	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileContent, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var jsonFileInterface interface{}
	err = UnmarshalJsonFile(fileContent, &jsonFileInterface)
	return jsonFileInterface, err
}

func UnmarshalJsonFile(jsonFile []byte, jsonFileInterface *interface{}) error {
	return json.Unmarshal(jsonFile, jsonFileInterface)
}
