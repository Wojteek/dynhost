package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Cloudflare - main structure of the Cloudflare
type Cloudflare struct {
	AuthToken      string
	Hostname       string
	IP             string
	ZoneIdentifier string
	DNSIdentifier  string
}

// CFUpdateRecordRequest - the structure uses for updating the record
type CFUpdateRecordRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
	Record  string `json:"type"`
	Proxied bool   `json:"proxied"`
}

// CFErrors - structure of the error response
type CFErrors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// CFRequestUpdatePayload structure of update request
type CFRequestUpdatePayload struct {
	Success bool       `json:"success"`
	Errors  []CFErrors `json:"errors"`
}

const (
	// CfUpdateEndpoint uses for updating the dns record
	CfUpdateEndpoint = "https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s"
)

// NewCloudflare is a structure of the Cloudflare
func NewCloudflare(
	authToken string,
	hostname string,
	ip string,
	zoneIdentifier string,
	dnsIdentifier string,
) *Cloudflare {
	c := &Cloudflare{
		AuthToken:      authToken,
		Hostname:       hostname,
		IP:             ip,
		ZoneIdentifier: zoneIdentifier,
		DNSIdentifier:  dnsIdentifier,
	}

	return c
}

// UpdateRecordRequest updates the DNS records in the Cloudflare
func (c *Cloudflare) UpdateRecordRequest() ([]byte, error) {
	data, err := c.prepareDataRequest()

	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest(data)
	var payload CFRequestUpdatePayload

	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}

	if !payload.Success {
		err, _ := json.Marshal(payload.Errors)

		return nil, fmt.Errorf("cloudflare error: %s", string(err))
	}

	return body, err
}

func (c *Cloudflare) sendRequest(data []byte) ([]byte, error) {
	req, err := c.newRequest(data)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	r, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(r io.ReadCloser) {
		if errReq := r.Close(); errReq != nil {
			err = errReq
		}
	}(r.Body)

	return ioutil.ReadAll(r.Body)
}

func (c *Cloudflare) newRequest(data []byte) (*http.Request, error) {
	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf(CfUpdateEndpoint, c.ZoneIdentifier, c.DNSIdentifier),
		bytes.NewBuffer(data),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))

	return req, nil
}

func (c *Cloudflare) prepareDataRequest() ([]byte, error) {
	data := CFUpdateRecordRequest{
		Name:    c.Hostname,
		Content: c.IP,
		Record:  "A",
		TTL:     1,
		Proxied: false,
	}

	return json.Marshal(data)
}
