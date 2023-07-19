package workout

import (
	"fmt"
	"time"
)

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

// func main() {
// 	currWeeksTimestamp := getWeeksTimeStamp(0)
// 	fmt.Printf("Current Week's timestamp: %v\n", currWeeksTimestamp)
// 	lastWeeksTimestamp := getWeeksTimeStamp(1)
// 	fmt.Printf("Previous Week's timestamp: %v\n", lastWeeksTimestamp)
// }
