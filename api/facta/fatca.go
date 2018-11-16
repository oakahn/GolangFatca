package facta

import (
	"encoding/xml"
	"log"

	"github.com/gin-gonic/gin"
	API "github.com/oakahn/GolangFatca/api"
	Data "github.com/oakahn/GolangFatca/data"
	Request "github.com/oakahn/GolangFatca/model/request"
	Response "github.com/oakahn/GolangFatca/model/response"
	Url "github.com/oakahn/GolangFatca/web"
)

func CallFatca() gin.HandlerFunc {
	// datas := getData()
	var requests Request.Envelope

	return func(c *gin.Context) {

		err := c.Bind(&requests)

		if err != nil {
			c.AbortWithStatusJSON(500, map[string]string{"err": err.Error()})
			return
		}

		c.JSON(200, requests)
	}
}

var tests = `{
	"soapenv:Envelope": {
	  "-xmlns:soapenv": "http://schemas.xmlsoap.org/soap/envelope/",
	  "-xmlns:ejbs": "http://ejbs",
	  "soapenv:Body": {
		"ejbs:getPartyFATCAInfo": {
		  "request": {
			"control": {
			  "branch": "00000",
			  "channel": "Test",
			  "requestId": "20160125102700",
			  "requesterName": "Test",
			  "user": "585190"
			},
			"customerId": "79221",
			"customerSource": "CBS"
		  }
		}
	  }
	}
  }`

func getData() Response.Envelope {
	resp, _ := API.Post(Url.Facta, Data.MockData())

	superman := Response.Envelope{}

	jsonErr := xml.Unmarshal(resp, &superman)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return superman
}
