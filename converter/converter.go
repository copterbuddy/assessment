package converter

import (
	"encoding/json"
	"net/http/httptest"
)

func ReqString(reqStruct interface{}) string {
	if reqStruct == nil {
		return ""
	}

	result, err := json.Marshal(&reqStruct)
	if err != nil {
		return ""
	}
	return string(result)
}

func ResStruct(res *httptest.ResponseRecorder, result interface{}) error {
	return json.Unmarshal([]byte(res.Body.Bytes()), &result)
}
