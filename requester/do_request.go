package requester

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

func (r *Request) DoRequest() error {
	start := time.Now()

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: customTransport}

	res, err := client.Do(r.internalRequestToFire)
	r.TimeNeeded = time.Since(start)

	if err != nil {
		return err
	}

	r.Response = res

	all, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	r.ResponseBody = string(all)

	return nil
}
