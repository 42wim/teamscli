package token

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/42wim/sqlittle"
)

var (
	ErrToken = errors.New("retrieving token failed")
)

func Read() (string, error) {
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
