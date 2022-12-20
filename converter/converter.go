package converter

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
)

func ReqString(reqStruct interface{}) *strings.Reader {
	if reqStruct != nil {
		return strings.NewReader("")
	}
	result, _ := json.Marshal(&reqStruct)
	return strings.NewReader(string(result))
}

func ResStruct(res *httptest.ResponseRecorder, result interface{}) {
	json.Unmarshal([]byte(res.Body.Bytes()), &result)
}
