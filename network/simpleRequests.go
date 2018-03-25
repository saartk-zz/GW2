package network

import (
	"net/http"
	"time"
	"fmt"
	"errors"
)

type simpleWebClient struct {
	client *http.Client
}


func GetClient (timeoutSecs ...int) (*simpleWebClient,error) {
	c := &http.Client{}
	if len(timeoutSecs)> 0 {
		duration,err := time.ParseDuration(fmt.Sprintf("%ss",timeoutSecs[0]))
		if err != nil {
			return nil,errors.New(fmt.Sprintf("Could not create http client. Error: %v",err))
		}
		c.Timeout = duration
	}
	return &simpleWebClient{client:c},nil
}

func GetJson(){

}
