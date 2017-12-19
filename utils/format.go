package utils

import (
	"encoding/json"
	"fmt"
)

func Convert2Json(i interface{}) string {
	var res string
	jsonResponse, err := json.Marshal(i)
	if err != nil {
		// todo logot irj ide
		res = fmt.Sprintf("%#v", i)
	} else {
		res = fmt.Sprintf("%s", jsonResponse)
	}
	return res
}
