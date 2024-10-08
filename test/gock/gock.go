package gock

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ReqParam struct {
	X int `json:"x"`
}

type Result struct {
	Value int `json:"value"`
}

func GetResultByAPI(x, y int) int {
	p := &ReqParam{X: x}
	b, _ := json.Marshal(p)

	resp, err := http.Post("http://your-api.com/post", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return -1
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1
	}
	var ret Result
	if err := json.Unmarshal(body, &ret); err != nil {
		return -1
	}
	return ret.Value + y
}
