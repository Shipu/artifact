package artifact

import "encoding/json"

func toJSON(v interface{}) string {
	bytes, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic("parse json error")
	}
	return string(bytes)
}
