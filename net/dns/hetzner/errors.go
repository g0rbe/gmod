package hetzner

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNoAPIKey          = errors.New("no API key found in request")
	ErrInvalidAPIKey     = errors.New("invalid authentication credentials")
	ErrZoneNotFound      = errors.New("zone not found")
	ErrInvalidArgument   = errors.New("invalid argument")
	ErrInvalidARecord    = errors.New("invalid A record")
	ErrInvalidAAAARecord = errors.New("invalid AAAA record")
)

func parseError(code int, body []byte) error {

	// 401 has a different returned body.
	// eg.: {"message":"Invalid authentication credentials"}
	if code == http.StatusUnauthorized {

		v := struct {
			Message string `json:"message"`
		}{}

		err := json.Unmarshal(body, &v)
		if err != nil {
			return fmt.Errorf("failed to unmarshal error body: %w", err)
		}

		switch v.Message {
		case "No API key found in request":
			return ErrNoAPIKey
		case "Invalid authentication credentials":
			return ErrInvalidAPIKey
		default:
			return fmt.Errorf("%s (%d)", v.Message, code)
		}

	}

	// If error occured, the error struct is inside the returned struct
	// eg.: {"zones": ..., "meta": ..., "error": ...}

	// This strict is unmarshal the error field only
	v := struct {
		Error struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		} `json:"error"`
	}{}

	err := json.Unmarshal(body, &v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal error body: %w, body: %s", err, body)
	}

	// Compare errors to known errors
	switch v.Error.Message {
	case "zone not found":
		return ErrZoneNotFound
	case "invalid A record":
		return ErrInvalidARecord
	case "invalid AAAA record":
		return ErrInvalidAAAARecord
	default:
		return fmt.Errorf("%s (%d)", v.Error.Message, code)
	}
}
