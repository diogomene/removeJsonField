package util

import (
	"github.com/diogomene/removeJsonField/entities"
)

func RemoveFieldsFromJsonInterface(jsonFile *entities.JsonFile, fieldsToRemove []string) {

	switch (*jsonFile).JsonContentType {
	case entities.ARRAY:
		if contentArray, ok := (*(*jsonFile).Content).([]interface{}); ok {
			for _, item := range contentArray {
				if contentMap, ok := item.(map[string]interface{}); ok {
					for _, fieldToRemove := range fieldsToRemove {
						delete(contentMap, fieldToRemove)
					}
				}
			}
		}
	case entities.OBJECT:
		if contentMap, ok := (*(*jsonFile).Content).(map[string]interface{}); ok {
			for _, fieldToRemove := range fieldsToRemove {
				delete(contentMap, fieldToRemove)
			}
		}
	}
}
