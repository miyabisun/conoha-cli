package endpoints

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(url string, tokenId string, reqBody string, res *Response) error {
	client := &http.Client{}
	buffer := bytes.NewBufferString(reqBody)
	request, err := http.NewRequest("POST", url, buffer)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
		return errors.New(fmt.Sprintf("POST to %s: failed (status code: %s).", url, res.Status))
	}
	body, err := ioutil.ReadAll(resp.Body)
	res.Body = body
	return err
}
