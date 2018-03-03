package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

func create(method string, url *url.URL, body interface{}, header map[string]string) (*http.Request, error) {
	bByte, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bReader := bytes.NewReader(bByte)
	u := url.String()
	req, err := http.NewRequest(method, u, bReader)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	return req, nil
}

func do(request *http.Request, response interface{}) error {
	c := &http.Client{}
	r, err := c.Do(request)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	if r.StatusCode != 200 {
		return NewResponseError(r.StatusCode, request, string(b))
	}
	if response == nil {
		return nil
	}
	return json.Unmarshal(b, response)
}

func createURL(endpoint string, relPath string, params map[string][]string) (*url.URL, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, relPath)
	if params == nil {
		return u, nil
	}
	v := url.Values(params)
	u.RawQuery = v.Encode()
	return u, nil
}

//Post Post
func Post(endpoint string, relPath string, params map[string][]string, body interface{}, header map[string]string, response interface{}) error {
	u, err := createURL(endpoint, relPath, params)
	if err != nil {
		return err
	}
	r, err := create("POST", u, body, header)
	if err != nil {
		return err
	}
	return do(r, response)
}
