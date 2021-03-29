package status

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/42wim/teamscli/internal/token"
)

var StatusURL = "https://presence.teams.microsoft.com/v1/me/publishnote"

type request struct {
	Expiry  string `json:"expiry,omitempty"`
	Message string `json:"message"`
}

func setStatus(status, duration string, showonmessage bool, token string) error {
	var expiry string

	bearer := "Bearer " + token

	if duration != "" {
		d, err := time.ParseDuration(duration)
		if err != nil {
			return err
		}

		expiry = time.Now().Add(d).UTC().Format(time.RFC3339Nano)
	} else {
		t, _ := time.Parse(time.RFC3339Nano, "9999-12-30T23:00:00.000Z")
		expiry = t.UTC().Format(time.RFC3339Nano)
	}

	if showonmessage {
		status += "<pinnednote></pinnednote>"
	}

	bReq, err := json.Marshal(&request{
		Expiry:  expiry,
		Message: status,
	})
	if err != nil {
		return fmt.Errorf("json marshal failed: %s", err)
	}

	req, _ := http.NewRequest("PUT", StatusURL, bytes.NewReader(bReq))

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

func SetStatusDuration(status, duration string, showonmessage bool) error {
	token, err := token.Read()
	if err != nil {
		return err
	}

	return setStatus(status, duration, showonmessage, token)
}

func SetStatus(status string, showonmessage bool) error {
	return SetStatusDuration(status, "", showonmessage)
}
