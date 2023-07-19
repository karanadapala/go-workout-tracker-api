package main

import (
	"goworkouttrackerapi/workout"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Timestamp struct {
	Timestamp string `json:"timestamp"`
}

func getCurrWeeksTimestamp(context *gin.Context) {
	timestamp := workout.GetWeeksTimeStamp(0)
	returnObj := Timestamp{
		Timestamp: timestamp,
	}
	context.IndentedJSON(http.StatusOK, returnObj)
}
func getPrevWeeksTimestamp(context *gin.Context) {
	timestamp := workout.GetWeeksTimeStamp(1)
	returnObj := Timestamp{
		Timestamp: timestamp,
	}
	context.IndentedJSON(http.StatusOK, returnObj)
}

func main() {
	router := gin.Default()
	router.GET("/this-week", getCurrWeeksTimestamp)
	router.GET("/prev-week", getPrevWeeksTimestamp)
	router.Run("localhost:8989")
}
