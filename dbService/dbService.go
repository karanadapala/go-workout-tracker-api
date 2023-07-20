package dbservice

import (
	"database/sql"
	"encoding/json"
	"goworkouttrackerapi/config"
	"goworkouttrackerapi/errors"

	_ "github.com/mattn/go-sqlite3"
)

const DB_PATH = "../local.db"

func ConnectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", DB_PATH)
	errors.CheckErr(err)
	return db
}

func InitailizeDB(db *sql.DB) {
	for _, v := range config.InitDBs {
		_, err := db.Exec(v)
		errors.CheckErr(err)
	}
}

func CloseDB(db *sql.DB) {
	db.Close()
}

func GetAllRowsFromWorkout(db *sql.DB) []WorkoutTable {
	rows, err := db.Query("SELECT * FROM workouts")
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
