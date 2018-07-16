package httpCalls

import (
	"net/http"
	"time"
	"io/ioutil"

	"github.com/saartk/GW2.git/network/consts"
)

type httpCaller struct {
	client *http.Client
}

func NewHttpSimpleCaller(timeout time.Duration) HttpCaller {
	client := &http.Client{
		Timeout:timeout,
	}
	if timeout == 0 {
		client.Timeout = consts.DefaultHttpClientTimeout
	}
	return &httpCaller{client:client}
}

func(h *httpCaller) Get(path string) ([]byte,error) {
	resp,err := h.client.Get(path)
	if err != nil {
		return nil,err
	}
	return ioutil.ReadAll(resp.Body)
}
