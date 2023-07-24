package dbservice

import (
	"database/sql"
	"encoding/json"
	"goworkouttrackerapi/config"
	"goworkouttrackerapi/errors"

	_ "github.com/mattn/go-sqlite3"
)

func ExecuteSelectQueryInWorkoutsDB(db *sql.DB, query string) []WorkoutTable {
	rows, err := db.Query(query)
	errors.CheckErr(err)

	var weekTimestamp string
	var dayTimestamp string
	var stringifiedWorkoutPlan []byte
	var workouts []WorkoutTable
	for rows.Next() {
		err := rows.Scan(&weekTimestamp, dayTimestamp, stringifiedWorkoutPlan)
		errors.CheckErr(err)
		var dailyPlan config.DailyPlan

		json.Unmarshal(stringifiedWorkoutPlan, &dailyPlan)

		dayPlan := WorkoutTable{
			WeekTimestamp: weekTimestamp,
			DayTimestamp:  dayTimestamp,
			WorkoutPlan:   dailyPlan,
		}
		workouts = append(workouts, dayPlan)
	}

	return workouts
}

func GetWeeksWorkout(timestamp string) []WorkoutTable {
	db := ConnectToDB()
	InitailizeDB(db)
	weekWorkouts := ExecuteSelectQueryInWorkoutsDB(db, timestamp)
	CloseDB(db)

	return weekWorkouts
}
