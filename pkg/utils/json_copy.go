package utils

import "encoding/json"

func CopyJsonStruct(data, dest interface{}) {
	jsonByte, _ := json.Marshal(data)
	json.Unmarshal(jsonByte, &dest)
}
