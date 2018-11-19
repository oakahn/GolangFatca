package request

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Ejbs    string   `xml:"xmlns:ejbs,attr"`
	Header  Header   `xml:"soapenv:Header"`
	Body    Body     `xml:"soapenv:Body"`
}

type Header struct {
}

type Body struct {
	GetPartyFATCAInfo GetPartyFATCAInfo `xml:"ejbs:getPartyFATCAInfo"`
}

type GetPartyFATCAInfo struct {
	Request Request `xml:"request" json:"request"`
}

type Request struct {
	Control        Control `xml:"control"`
	CustomerID     string  `xml:"customerId"`
	CustomerSource string  `xml:"customerSource"`
}

type Control struct {
	Branch        string `xml:"branch"`
	Channel       string `xml:"channel"`
	RequestId     string `xml:"requestId"`
	RequesterName string `xml:"requesterName"`
	User          string `xml:"user"`
}
