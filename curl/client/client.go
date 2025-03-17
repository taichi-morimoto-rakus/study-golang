package client

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type HttpClient struct {
	url           *url.URL
	method        string
	requestBody   *string
	requestHeader map[string]string
}

func NewHttpClient(
	rawurl string,
	method string,
	data string,
	customHeaders []string,
) (*HttpClient, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	requestHeader := make(map[string]string)
	for _, header := range customHeaders {
		kv := strings.Split(header, ": ")
		requestHeader[kv[0]] = kv[1]
	}

	var requestBody *string
	if method == http.MethodGet || method == http.MethodDelete {
		requestBody = nil
		delete(requestHeader, "Content-Type")
	} else {
		if len(data) == 0 {
			return nil, errors.New("requests data json")
		}
		requestBody = &data
		requestHeader["Content-Type"] = "application/json"
	}

	return &HttpClient{
		url:           url,
		method:        method,
		requestBody:   requestBody,
		requestHeader: requestHeader,
	}, nil
}

func (c *HttpClient) Execute() (string, string, error) {
	req, res, err := c.SendRequest()
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	return CreateRequestText(req), CreateResponseText(res), nil
}

func (c *HttpClient) SendRequest() (*http.Request, *http.Response, error) {
	var body io.Reader
	if c.requestBody != nil {
		body = strings.NewReader(*c.requestBody)
	}
	req, err := http.NewRequest(
		c.method,
		c.url.String(),
		body,
	)
	if err != nil {
		return nil, nil, err
	}

	for k, v := range c.requestHeader {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	return req, res, nil
}

func CreateRequestText(req *http.Request) string {
	output := "\n===Request===\n"
	output += "[URL] " + req.URL.String() + "\n"
	output += "[Method] " + req.Method + "\n"
	output += "[Headers]\n"
	for _, k := range sortedKeys(req.Header) {
		output += "  " + k + ": " + strings.Join(req.Header[k], "; ") + "\n"
	}
	return output
}

func CreateResponseText(res *http.Response) string {
	output := "\n===Response===\n"
	output += "[Status] " + strconv.Itoa(res.StatusCode) + "\n"
	output += "[Headers]\n"
	for _, k := range sortedKeys(res.Header) {
		output += "  " + k + ": " + strings.Join(res.Header[k], "; ") + "\n"
	}
	output += "[Body]\n"
	bodyBytes, _ := io.ReadAll(res.Body)
	output += string(bodyBytes) + "\n"
	return output
}

// http.Request.Header と http.Response.Header を渡すと昇順にソートされた Key を返す関数
func sortedKeys(m map[string][]string) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}

	sort.Strings(s)
	return s
}
