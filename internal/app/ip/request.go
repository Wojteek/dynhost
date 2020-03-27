package ip

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func request(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
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
