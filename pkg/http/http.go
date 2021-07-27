package http

import (
	"io/ioutil"
	"net/http"
)

type HttpResponse struct {
	StatusCode	int
	Body		[]byte
}

func DoGet(url string) (*HttpResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &HttpResponse{Body: body, StatusCode: resp.StatusCode}
	return response,nil
}