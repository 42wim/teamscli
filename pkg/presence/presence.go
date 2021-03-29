package presence

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/42wim/teamscli/internal/client"
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

	return client.PUT(PresenceURL, bReq, token)
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
