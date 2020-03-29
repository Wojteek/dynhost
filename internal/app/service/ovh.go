package service

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
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
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", o.Credentials.Username, o.Credentials.Password)))))
	r, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(r io.ReadCloser) {
		if errReq := r.Close(); errReq != nil {
			err = errReq
		}
	}(r.Body)

	if err := r.Header.Get("WWW-Authenticate"); r.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("OVH [%d]: %s", r.StatusCode, err)
	}

	resp, err := ioutil.ReadAll(r.Body)

	if re := regexp.MustCompile(`nochg .*?`); !re.MatchString(string(resp)) {
		return nil, fmt.Errorf("OVH response: %s", resp)
	}

	return resp, err
}
