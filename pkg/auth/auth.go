package auth

import (
	"encoding/base64"
)

type Auth interface {
	Header() map[string]string
}

type BasicAuth struct {
	Username string
	Password string
}

func (b BasicAuth) Header() map[string]string {
	authValue := base64.StdEncoding.EncodeToString([]byte(b.Username + ":" + b.Password))

	return map[string]string{
		"Authorization": "Basic " + authValue,
	}
}

type APIKeyAuth struct {
	Key        string
	HeaderName string
}

func (a APIKeyAuth) Header() map[string]string {
	return map[string]string{
		a.HeaderName: a.Key,
	}
}
