package hetzner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TXTVerification struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Pagination struct {
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	LastPage     int `json:"last_page"`
	TotalEntries int `json:"total_entries"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Zone struct {
	ID              string          `json:"id"`
	Created         string          `json:"created"`
	Modified        string          `json:"modified"`
	LegacyDNSHost   string          `json:"legacy_dns_host"`
	LegacyNS        []string        `json:"legacy_ns"`
	Name            string          `json:"name"`
	NS              []string        `json:"ns"`
	Owner           string          `json:"owner"`
	Paused          bool            `json:"paused"`
	Permission      string          `json:"permission"`
	Project         string          `json:"project"`
	Registrar       string          `json:"registrar"`
	Status          string          `json:"status"`
	TTL             int             `json:"ttl"`
	Verified        string          `json:"verified"`
	RecordsCount    int             `json:"records_count"`
	IsSecondaryDNS  bool            `json:"is_secondary_dns"`
	TXTVerification TXTVerification `json:"txt_verification"`
}

type Zones struct {
	Zones []Zone `json:"zones"`
	Meta  Meta   `json:"meta"`
}

// GetAllZones returns every zones associated with the user.
func (c *Client) GetAllZones() ([]Zone, error) {

	req, err := http.NewRequest("GET", BaseURL+"/zones", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Auth-API-Token", c.key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed request: %w", err)
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Error
	if resp.StatusCode != http.StatusOK {

		return nil, parseError(resp.StatusCode, respBody)
	}

	zones := new(Zones)

	err = json.Unmarshal(respBody, zones)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return zones.Zones, nil
}

// GetZoneByName returns the zone associated with the user with name name.
func (c *Client) GetZoneByName(name string) (Zone, error) {

	req, err := http.NewRequest("GET", BaseURL+"/zones?name="+name, nil)
	if err != nil {
		return Zone{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Auth-API-Token", c.key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Zone{}, fmt.Errorf("failed request: %w", err)
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Zone{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Error
	if resp.StatusCode != http.StatusOK {

		return Zone{}, parseError(resp.StatusCode, respBody)
	}

	zones := new(Zones)

	err = json.Unmarshal(respBody, zones)
	if err != nil {
		return Zone{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(zones.Zones) < 1 {
		return Zone{}, ErrZoneNotFound
	}

	if len(zones.Zones) > 1 {
		return Zone{}, fmt.Errorf("multiple zone returned: %d", len(zones.Zones))
	}

	return zones.Zones[0], nil
}
