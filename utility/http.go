package utility

import (
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

type Response struct {
	statusCode int
	body       []byte
}

func (resp Response) Is200() bool {
	return resp.statusCode == http.StatusOK
}

func (resp Response) Body() []byte {
	return resp.body
}

func Do(req http.Request) (Response, error) {
	resp, err := httpClient.Do(&req)
	if err != nil {
		return Response{}, err
	}
	if resp.ContentLength == 0 {
		return Response{
			statusCode: resp.StatusCode,
		}, nil
	}
	defer resp.Body.Close()
	bytes, err := ReadBytes(resp.Body)
	if err != nil {
		return Response{}, err
	}
	return Response{
		statusCode: resp.StatusCode,
		body:       bytes,
	}, nil
}
