package hetzner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Record struct {
	Type     string
	ID       string
	Created  string
	Modified string
	ZoneID   string `json:"zone_id"`
	Name     string
	Value    string
	TTL      int
}

type Records struct {
	Records []Record
}

// GetAllRecords returns all records associated with user.
func (c *Client) GetAllRecords() ([]Record, error) {

	req, err := http.NewRequest("GET", BaseURL+"/records", nil)
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

	v := new(Records)

	err = json.Unmarshal(respBody, v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return v.Records, nil
}

// GetAllRecordsByZone returns all records associated with user from zone zone.
func (c *Client) GetAllRecordsByZone(zone string) ([]Record, error) {

	req, err := http.NewRequest("GET", BaseURL+"/records?zone_id="+zone, nil)
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

	v := new(Records)

	err = json.Unmarshal(respBody, v)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return v.Records, nil
}

// CreateRecord creates a new record.
// Valid t types are: "A", "AAAA", "NS", "MX", "CNAME", "RP", "TXT", "SOA", "HINFO", "SRV", "DANE", "TLSA", "DS" and "CAA".
func (c *Client) CreateRecord(name string, ttl int, t string, value string, zone string) (Record, error) {

	body := fmt.Sprintf("{\"name\":\"%s\",\"ttl\":%d,\"type\":\"%s\",\"value\":\"%s\",\"zone_id\":\"%s\"}", name, ttl, t, value, zone)

	req, err := http.NewRequest("POST", BaseURL+"/records", strings.NewReader(body))
	if err != nil {
		return Record{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Auth-API-Token", c.key)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Record{}, fmt.Errorf("failed request: %w", err)
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Record{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Error
	if resp.StatusCode != http.StatusOK {

		return Record{}, parseError(resp.StatusCode, respBody)
	}

	v := struct {
		Record Record `json:"record"`
	}{}

	err = json.Unmarshal(respBody, &v)
	if err != nil {
		return Record{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return v.Record, nil
}

// DeleteRecord deletes a record with id id.
func (c *Client) DeleteRecord(id string) error {

	req, err := http.NewRequest("DELETE", BaseURL+"/records/"+id, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Auth-API-Token", c.key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	// Read Response Body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	return parseError(resp.StatusCode, respBody)
}
