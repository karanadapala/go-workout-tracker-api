package workout

import (
	"fmt"
	"goworkouttrackerapi/config"
	"goworkouttrackerapi/errors"
	"time"
)

var weekdayMap = map[config.TDay]time.Weekday{
	"sunday":    time.Sunday,
	"monday":    time.Monday,
	"tuesday":   time.Tuesday,
	"wednesday": time.Wednesday,
	"thursday":  time.Thursday,
	"friday":    time.Friday,
	"saturday":  time.Saturday,
}

func GetWeeksTimeStamp(weekOffSet int) string {
	currDayTimestamp := time.Now().Truncate(time.Hour * 24)
	fmt.Printf("Current Day's Timestamp: %v\n", currDayTimestamp)
	currWeekDay := currDayTimestamp.Weekday()
	fmt.Printf("Current Week Day: %v\n", currWeekDay)
	sinceWeekStart := time.Duration(currWeekDay)*24*time.Hour + time.Duration(currDayTimestamp.Hour()*int(time.Hour))
	fmt.Printf("Hours since Week's start: %v\n", -sinceWeekStart)
	currWeeksTimestamp := currDayTimestamp.Add(-sinceWeekStart)
	offsetTimestamp := currWeeksTimestamp.Add(-(time.Duration(weekOffSet) * 24 * 7 * time.Hour))
	return offsetTimestamp.Format(time.RFC3339)
}

func GetDayTimestamp(day config.TDay, weekrfc3339String string) string {
	weeksTime, err := time.Parse(time.RFC3339, weekrfc3339String)
	errors.CheckErr(err)
	weekDayTime := time.Duration(weekdayMap[day]) * 24 * time.Hour
	daysTime := weeksTime.Add((weekDayTime))
	return daysTime.Format(time.RFC3339)
}
