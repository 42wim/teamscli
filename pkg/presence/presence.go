package presence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/42wim/sqlittle"
)

var (
	ErrToken  = errors.New("retrieving token failed")
	StatusURL = "https://presence.teams.microsoft.com/v1/me/forceavailability/"
)

type request struct {
	Availability          string `json:"availability"`
	Activity              string `json:"activity,omitempty"`
	DesiredExpirationTime string `json:"desiredExpirationTime,omitempty"`
}

func readToken() (string, error) {
	var token string

	configfile, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "windows":
		configfile = filepath.Join(configfile, "Microsoft", "Teams", "Cookies")
	case "linux":
		configfile = filepath.Join(configfile, "Microsoft", "Microsoft Teams", "Cookies")
	}

	db, err := sqlittle.Open(configfile)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	defer db.Close()

	// read the TSAUTHCOOKIE from the cookies table
	err = db.Select("cookies", func(r sqlittle.Row) {
		var (
			name  string
			value string
		)

		err = r.Scan(&name, &value)
		if err != nil {
			log.Panicf("db scan failed: %s", err)
		}

		if name == "TSAUTHCOOKIE" {
			token = value
		}
	}, "name", "value")
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrToken, err)
	}

	return token, nil
}

func setPresence(presence, token string) error {
	var activity string

	bearer := "Bearer " + token

	if strings.ToLower(presence) == "offline" {
		activity = "OffWork"
	}

	bReq, err := json.Marshal(&request{
		Availability: presence,
		Activity:     activity,
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

func SetPresence(presence string) error {
	token, err := readToken()
	if err != nil {
		return err
	}

	return setPresence(presence, token)
}
