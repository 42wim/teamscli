package presence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/42wim/teamscli/internal/token"
)

var (
	PresenceURL = "https://presence.teams.microsoft.com/v1/me/forceavailability/"
)

type request struct {
	Availability          string `json:"availability"`
	Activity              string `json:"activity,omitempty"`
	DesiredExpirationTime string `json:"desiredExpirationTime,omitempty"`
}

func setPresence(presence, duration, token string) error {
	var activity string
	var desiredExpirationTime string

	bearer := "Bearer " + token

	if duration != "" {
		d, err := time.ParseDuration(duration)
		if err != nil {
			return err
		}

		desiredExpirationTime = time.Now().Add(d).UTC().Format(time.RFC3339Nano)
	}

	if strings.ToLower(presence) == "offline" {
		activity = "OffWork"
	}

	bReq, err := json.Marshal(&request{
		Availability:          presence,
		Activity:              activity,
		DesiredExpirationTime: desiredExpirationTime,
	})
	if err != nil {
		return fmt.Errorf("json marshal failed: %s", err)
	}

	req, _ := http.NewRequest("PUT", PresenceURL, bytes.NewReader(bReq))

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

func SetPresenceDuration(presence, duration string) error {
	token, err := token.Read()
	if err != nil {
		return err
	}

	switch strings.ToLower(presence) {
	case "dnd":
		presence = "DoNotDisturb"
	case "brb":
		presence = "BeRightBack"
	}

	return setPresence(presence, duration, token)
}

func SetPresence(presence string) error {
	return SetPresenceDuration(presence, "")
}
