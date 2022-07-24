package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Client ...
type Client struct {
	URL        string
	HTTPClient *http.Client
}

// Response ...
type Response struct {
	StatusCode int
	Body       *[]byte
}

// CreateClient ...
func CreateClient(url string) *Client {
	client := &http.Client{}
	newClient := &Client{url, client}
	return newClient
}

// MakeRequest ...
func (c *Client) MakeRequest(method string, reqBody *[]byte, path string) (Response, error) {
	req, err := http.NewRequest(method, c.URL+path, bytes.NewReader(*reqBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	clientResponse := Response{
		resp.StatusCode,
		&respBody,
	}
	return clientResponse, nil
}
