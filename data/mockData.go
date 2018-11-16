package data

import (
	"encoding/xml"
	"log"

	Request "github.com/oakahn/GolangFatca/model/request"
)

func MockData() string {
	data := Request.Control{
		Branch:        "00000",
		Channel:       "Test",
		RequesterName: "Test",
		RequestId:     "20160125102700",
		User:          "585190",
	}

	request := Request.Request{
		Control:        data,
		CustomerID:     "79221",
		CustomerSource: "CBS",
	}

	getPartyFATCAInfo := Request.GetPartyFATCAInfo{
		Request: request,
	}

	body := Request.Body{
		GetPartyFATCAInfo: getPartyFATCAInfo,
	}

	soapenvEnvelope := Request.Envelope{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		Ejbs:    "http://ejbs",
		Body:    body,
	}

	resp, err := xml.MarshalIndent(soapenvEnvelope, " ", "  ")

	if err != nil {
		log.Fatal(err)
	}

	return string(resp)
}
