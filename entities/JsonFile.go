package entities

type JsonContentType int

const (
	ARRAY JsonContentType = iota
	OBJECT
	UNDEFINED
)

type JsonFile struct {
	Content         *interface{}
	Fields          []string
	JsonContentType JsonContentType
}

func NewJsonFile(jsonInterface *interface{}) *JsonFile {
	contentType, fields := GetJsonFileMeta(*jsonInterface)
	return &JsonFile{
		Content:         jsonInterface,
		Fields:          fields,
		JsonContentType: contentType,
	}
}

func GetJsonFileMeta(json interface{}) (JsonContentType, []string) {
	jsonContentType := UNDEFINED
	var mapArray []map[string]interface{}
	switch json := json.(type) {
	case map[string]interface{}:
		jsonContentType = OBJECT
		mapArray = []map[string]interface{}{json}
	case []interface{}:
		jsonContentType = ARRAY
		for _, item := range json {
			mapArray = append(mapArray, item.(map[string]interface{}))
		}
	}

	mockMap := make(map[string]byte)
	for _, item := range mapArray {
		for key := range item {
			mockMap[key] = 0x0
		}
	}

	var fields []string

	for key := range mockMap {
		fields = append(fields, key)
	}

	return jsonContentType, fields
}
