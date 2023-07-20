package dbservice

import "goworkouttrackerapi/config"

type DB struct {
	Workout WorkoutTable
}

type WorkoutTable struct {
	WeekTimestamp string
	DayTimestamp  string
	WorkoutPlan   config.DailyPlan
}

type UserTable struct {
	WeekTimestamp string
	DayTimestamp  string
}
