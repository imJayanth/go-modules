package helpers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpRequest struct {
	Method        string            `json:"method"`
	Headers       map[string]string `json:"headers"`
	Url           string            `json:"url"`
	Body          interface{}       `json:"body"`
	Values        url.Values        `json:"values"`
	Authorization string            `json:"authorization"`
	ClientConfig  ClientConfig      `json:"client_config"`
}

type ClientConfig struct {
	InsecureSkipVerify bool `json:"insecure_skip_verify"`
}

type Response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Body       string `json:"body"`
	Error      error  `json:"error"`
}

func (request *HttpRequest) Do() *Response {
	reqBytes, e := json.Marshal(request.Body)
	if e != nil {
		return &Response{Error: e}
	}
	requestReader := bytes.NewReader(reqBytes)
	client := &http.Client{}
	if request.ClientConfig.InsecureSkipVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}
	req, e := http.NewRequest(request.Method, request.Url, requestReader)
	if e != nil {
		return &Response{Error: e}
	}
	req.Header.Set("Content-Type", "application/json")
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	if request.Authorization != "" {
		req.Header.Set("Authorization", request.Authorization)
	}
	resp, e := client.Do(req)
	if e != nil {
		return &Response{Error: e}
	}
	defer resp.Body.Close()
	respBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return &Response{Error: e}
	}
	return &Response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       string(respBody),
	}
}

func (request *HttpRequest) PostForm() *Response {
	resp, e := http.PostForm(request.Url, request.Values)
	if e != nil {
		return &Response{Error: e}
	}
	defer resp.Body.Close()
	respBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return &Response{Error: e}
	}
	return &Response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       string(respBody),
	}
}

func (response *Response) Unmarshal(data interface{}) error {
	if response.Error != nil {
		return response.Error
	}
	return json.Unmarshal([]byte(response.Body), &data)
}
