package httpCalls

import (
	"time"
	"io/ioutil"
	"net/http"
	"fmt"

	"github.com/saartk/GW2.git/network/consts"
)

type httpAuthCaller struct {
	client *http.Client
	auth   string
}

func NewHttpAuthCaller(timeout time.Duration, authKey string) HttpCaller {
	client := &http.Client{
		Timeout: timeout,
	}
	if timeout == 0 {
		client.Timeout = consts.DefaultHttpClientTimeout
	}
	return &httpAuthCaller{client: client, auth: authKey}
}

func (h *httpAuthCaller) Get(path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(consts.AuthHeaderKey, fmt.Sprintf(consts.AuthHeaderValue,h.auth))
	resp, err := h.client.Get(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
