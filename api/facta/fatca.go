package facta

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	API "github.com/oakahn/GolangFatca/api"
	Request "github.com/oakahn/GolangFatca/model/request"
	Response "github.com/oakahn/GolangFatca/model/response"
	Url "github.com/oakahn/GolangFatca/web"
)

func CallFatca() gin.HandlerFunc {

	return func(c *gin.Context) {

		var requests Request.Request

		err := c.BindJSON(&requests)

		dataXML := convertToXML(requests)

		if err != nil {
			c.AbortWithStatusJSON(500, map[string]string{"err": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, getData(dataXML).Body.GetPartyFATCAInfoResponse.GetPartyFATCAInfoReturn)
	}
}

func convertToXML(input Request.Request) string {
	control := input.Control
	data := Request.Control{
		Branch:        control.Branch,
		Channel:       control.Channel,
		RequesterName: control.RequesterName,
		RequestId:     control.RequestId,
		User:          control.User,
	}

	request := Request.Request{
		Control:        data,
		CustomerID:     input.CustomerID,
		CustomerSource: input.CustomerSource,
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

func getData(text string) Response.Envelope {

	// resp, _ := API.Post(Url.Facta, Data.MockData())
	resp, err := API.Post(Url.Facta, text)

	if err != nil {
		log.Fatal(err)
	}

	superman := Response.Envelope{}

	jsonErr := xml.Unmarshal(resp, &superman)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return superman
}
