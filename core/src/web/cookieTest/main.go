package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)


type deadlinedConn struct {
	Timeout time.Duration
	net.Conn
}

func NewDeadlineTransport(timeout time.Duration) *http.Transport {
	transport := &http.Transport{
		DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, timeout)
			if err != nil {
				return nil, err
			}
			return &deadlinedConn{timeout, c}, nil
		},
	}
	return transport
}


func HttpPostWithCookie(url string, cookies []*http.Cookie, headers map[string]string) ([]byte, int, error) {
	var (
		resp       *http.Response
		statusCode int
	)
	request, err := http.NewRequest("POST", url, nil)
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}
	httpclient := &http.Client{Transport: NewDeadlineTransport(5 * time.Second)}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	resp, err = httpclient.Do(request)

	if resp != nil {
		statusCode = resp.StatusCode
	}
	if err != nil {
		return nil, statusCode, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return b, statusCode, err
	}
	return b, statusCode, nil
}

func main() {
	url := "http://account-co.61.com/lvericode/generate"
	data, _, _ := HttpPostWithCookie(url, make([]*http.Cookie, 0 ,1) , make(map[string]string))
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Sprintf("数据： %v", str)
}

