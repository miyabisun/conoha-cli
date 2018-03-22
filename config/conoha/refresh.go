package conoha

import (
	"errors"
	"time"
)

func Refresh() error {
	var config Config
	err := Read(&config)
	if err != nil {
		return err
	}
	if isRequiredLogin(&config) {
		err = errors.New("required login.")
		return err
	}

	if isRequiredRelogin(&config) {
		err := relogin(&config)
		if err != nil {
			return err
		}
		err = Write(&config)
		if err != nil {
			return err
		}
	}
	return nil
}

func relogin(config *Config) error {
	return Login(&config.Auth, &config.Token)
}

func isRequiredLogin(config *Config) bool {
	return (config.Token.Id == "" || config.Token.Expires == "" || config.Auth.User == "" || config.Auth.Pass == "" || config.Auth.TenantId == "")
}

func isRequiredRelogin(config *Config) bool {
	format := "2006-01-02T15:04:05-0700"
	expires, err := time.Parse(format, config.Token.Expires)
	return (err != nil || expires.Unix() < time.Now().Unix())
}
