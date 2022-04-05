package app

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	ContentType string
	Url         string
	Body        any
}

func (r *Request) Post() (*http.Response, error) {
	body, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(r.Url, r.ContentType, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
