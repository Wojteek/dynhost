package provider

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// CredentialsOVH is a structure of credentials in OVH
type CredentialsOVH struct {
	Username string
	Password string
}

// OVH is a structure of updating DynHost OVH
type OVH struct {
	IP          string
	Hostname    string
	Credentials CredentialsOVH
}

// UpdateRecordRequest updates DynHost in OVH
func (o *OVH) UpdateRecordRequest() ([]byte, error) {
	client := &http.Client{}
	params := url.Values{
		"system":   {"dyndns"},
		"hostname": {o.Hostname},
		"ip":       {o.IP},
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.ovh.com/nic/update?%s", params.Encode()), nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", o.Credentials.Username, o.Credentials.Password)))))
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
