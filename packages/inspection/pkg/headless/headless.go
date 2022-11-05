package headless

import (
	"io"
	"io/ioutil"
	"net/http"
)

type VisitParams struct {
	Method  string
	Url     string
	Body    io.Reader
	Headers map[string]string
}

type HeadlessClient struct {
	client    *http.Client
	locations []string
}

type HeadlessVisitResult struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
	Locations  []string
}

func NewHeadlessClient() *HeadlessClient {
	c := &HeadlessClient{}
	c.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			c.locations = append(c.locations, req.URL.String())
			return nil
		},
	}
	return c
}

func (c *HeadlessClient) Visit(p VisitParams) (HeadlessVisitResult, error) {
	req, _ := http.NewRequest(p.Method, p.Url, p.Body)
	for k, v := range p.Headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return HeadlessVisitResult{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return HeadlessVisitResult{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
		Locations:  c.locations,
	}, nil
}
