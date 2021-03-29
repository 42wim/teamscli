package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func PUT(URL string, data []byte, token string) error {
	req, _ := http.NewRequest("PUT", URL, bytes.NewReader(data))
	bearer := "Bearer " + token

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ForceAttemptHTTP2:     true,
	}

	client := &http.Client{Transport: transport}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error while reading the response bytes: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("update failed: %s", string(body))
	}

	return nil
}
