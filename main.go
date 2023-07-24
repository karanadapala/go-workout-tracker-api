package main

import (
	"goworkouttrackerapi/config"
	dbservice "goworkouttrackerapi/dbService"
	"goworkouttrackerapi/errors"
	"goworkouttrackerapi/workout"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Timestamp struct {
	Timestamp string `json:"timestamp"`
}

type Workouts struct {
	Workouts []dbservice.WorkoutTable `json:"workouts"`
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

func getWorkouts(context *gin.Context) {
	db := dbservice.ConnectToDB()
	dbservice.InitailizeDB(db)
	workouts := dbservice.GetAllRowsFromWorkout(db)
	var returnObj = Workouts{
		Workouts: workouts,
	}
	context.IndentedJSON(http.StatusOK, returnObj)
}

func generateWorkout(context *gin.Context) {
	increaseVolumeParam := context.Query("increaseVolume")
	var increaseVolume bool
	if increaseVolumeParam != "" {
		_increaseVolume, err := strconv.ParseBool(increaseVolumeParam)
		increaseVolume = _increaseVolume
		errors.CheckErr(err)
	}
	weekPlan := workout.GenerateWeekPlan(config.BaseTemplate, increaseVolume)
	context.IndentedJSON(http.StatusOK, weekPlan)
}

func main() {
	router := gin.Default()
	router.GET("/this-week", getCurrWeeksTimestamp)
	router.GET("/prev-week", getPrevWeeksTimestamp)
	router.GET("/generate", generateWorkout)
	router.Run("localhost:8989")
}
