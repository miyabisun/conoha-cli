package endpoints

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(endpoint string, tokenId string, res *Response) error {
	client := &http.Client{}
	request, err := http.NewRequest("GET", endpoint, nil)
	request.Header.Add("Accept", "application/json")
	if tokenId != "" {
		request.Header.Add("X-Auth-Token", tokenId)
	}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res.Status = resp.StatusCode
	if res.Status < 200 || res.Status >= 300 {
		return errors.New(fmt.Sprintf("POST to %s: failed (status code: %s).", endpoint, res.Status))
	}
	body, err := ioutil.ReadAll(resp.Body)
	res.Body = body
	return err
}
