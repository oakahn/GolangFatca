package main

import (
	"github.com/gin-gonic/gin"
	Fatca "github.com/oakahn/GolangFatca/api/facta"
)

func main() {
	r := gin.Default()

	r.POST("/test", Fatca.CallFatca())

	r.Run(":1111")
}
