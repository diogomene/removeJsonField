package usecases

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/diogomene/removeJsonField/entities"
	"github.com/diogomene/removeJsonField/forms"
	"github.com/diogomene/removeJsonField/util"
)

func RemoveJsonFields() {
	var filePath string
	err := forms.GetJsonFilePath(&filePath)
	if err != nil {
		log.Fatal(err)
	}

	jsonFileInterface, err := util.ReadJsonFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	jsonFile := entities.NewJsonFile(&jsonFileInterface)
	fieldsToRemove, err := forms.DeleteJsonFieldsForm(*jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	util.RemoveFieldsFromJsonInterface(jsonFile, fieldsToRemove)
	outData, err := json.MarshalIndent(jsonFile.Content, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	outputFileName, err := forms.GetOutputFileName()
	if err != nil {
		log.Fatal(err)
	}
	outputFileName = strings.TrimSpace(outputFileName)
	if !strings.HasSuffix(outputFileName, ".json") {
		outputFileName += ".json"
	}
	os.WriteFile(outputFileName, outData, os.ModeAppend)
}
