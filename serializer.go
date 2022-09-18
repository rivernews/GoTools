package GoTools

import (
	"encoding/json"
)

func AsJson(v any) string {
	return string(AsJsonBytes(v))
}

func AsJsonBytes(v any) []byte {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		Logger("ERROR", err.Error())
	}
	return jsonBytes
}

func FromJson(b []byte, structInstance any) {
	if err := json.Unmarshal(b, structInstance); err != nil {
		Logger("ERROR", err.Error())
	}
}
