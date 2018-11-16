package api

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Post(url string, payload string) ([]byte, error) {

	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("SOAPAction", "getPartyFATCAInfo")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "498e204e-ee0e-4144-98a5-1a5f76cf5601")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
