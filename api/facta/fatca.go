package facta

import (
	"encoding/xml"
	"log"

	"github.com/gin-gonic/gin"
	API "github.com/oakahn/GolangFatca/api"
	Data "github.com/oakahn/GolangFatca/data"
	Response "github.com/oakahn/GolangFatca/model/Response"
	Url "github.com/oakahn/GolangFatca/web"
)

func CallFatca() gin.HandlerFunc {
	return func(c *gin.Context) {

		data := getData()

		c.JSON(200, gin.H{
			"Response": data,
		})
	}
}

func getData() Response.Envelope {
	resp, _ := API.Post(Url.Facta, Data.MockData())

	superman := Response.Envelope{}

	jsonErr := xml.Unmarshal(resp, &superman)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return superman
}
