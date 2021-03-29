package status

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/42wim/teamscli/internal/client"
	"github.com/42wim/teamscli/internal/token"
)

var StatusURL = "https://presence.teams.microsoft.com/v1/me/publishnote"

type request struct {
	Expiry  string `json:"expiry,omitempty"`
	Message string `json:"message"`
}

func setStatus(status, duration string, showonmessage bool, token string) error {
	var expiry string

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

	return client.PUT(StatusURL, bReq, token)
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
