package workout

import (
	"fmt"
	"goworkouttrackerapi/config"
	dbservice "goworkouttrackerapi/dbService"
	"goworkouttrackerapi/errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	// lop "github.com/samber/lo/parallel"
)

type Increase float64

const (
	TenPercent     Increase = 0.1
	TwentyPercent  Increase = 0.2
	ThirtyPercent  Increase = 0.3
	FortyPercent   Increase = 0.4
	FiftyPercent   Increase = 0.5
	SixtyPercent   Increase = 0.6
	SeventyPercent Increase = 0.7
	EightyPercent  Increase = 0.8
	NintyPercent   Increase = 0.9
	HundredPercent Increase = 1.0
)

func CalculateIncrease(number int, increase Increase) int {
	return int(math.Ceil(float64(number) * float64(increase)))
}

func GenerateWeekPlan(baseTemplate config.BaseWeeklyTemplate, increaseVolume bool) config.WeeklyPlan {
	generatedWeeklyBasePlan := make(config.WeeklyPlan)
	for day, baseWorkout := range config.BaseTemplate {
		generatedWeeklyBasePlan[day] = GenerateDayPlan(baseWorkout.MainWorkout, increaseVolume, day)
	}

	return generatedWeeklyBasePlan
}

func GenerateDayPlan(workoutIds []config.WorkoutNameIds, increaseVolume bool, day config.TDay) config.DailyPlan {
	var warmUps []config.WarmUpValue
	var coolDowns []config.CoolDownValue
	var accessoryWorkouts []config.DeterminedWorkout
	var mainWorkouts []config.DeterminedWorkout

	// Main Workouts
	mainWorkouts = lo.Map(workoutIds, func(workoutId config.WorkoutNameIds, index int) config.DeterminedWorkout {
		workoutDetails := ModifyExerciseDetails(workoutId, increaseVolume, day)
		return workoutDetails
	})

	var mainWorkoutsCategories []config.WorkoutTags

	if len(mainWorkouts) == 2 {
		mainWorkoutsCategories = mainWorkouts[0].Category
		mainWorkoutsCategories = append(mainWorkoutsCategories, mainWorkouts[1].Category...)
		mainWorkoutsCategories = lo.Uniq(mainWorkoutsCategories)
	}

	// Accessory Workouts
	accessoryWorkouts = GetAccessoryWorkouts(mainWorkoutsCategories, increaseVolume, day)

	//WarmUps

	dayPlan := config.DailyPlan{
		WarmUp:           warmUps,
		CoolDown:         coolDowns,
		AccessoryWorkout: accessoryWorkouts,
		MainWorkout:      mainWorkouts,
	}
	return dayPlan
}

func ModifyExerciseDetails(workoutId config.WorkoutNameIds, increaseVolume bool, day config.TDay) config.DeterminedWorkout {
	exerciseInfo := config.WorkoutNameConfig[workoutId]
	// Get sets information
	setsInfo := strings.Split(string(exerciseInfo.SetRange), "-")
	minSets, err := strconv.ParseInt(setsInfo[0], 10, 0)
	errors.CheckErr(err)
	maxSets, err := strconv.ParseInt(setsInfo[1], 10, 0)
	errors.CheckErr(err)

	// Get Reps Information
	repsInfo := strings.Split(string(exerciseInfo.RepRange), "-")
	minReps, err := strconv.Atoi(repsInfo[0])
	errors.CheckErr(err)
	maxReps, err := strconv.Atoi(repsInfo[1])
	errors.CheckErr(err)

	var newExercisePlan config.DeterminedWorkout

	prevWeek := GetWeeksTimeStamp(1)
	prevWeekSqlStatement := fmt.Sprintf("SELECT * FROM workouts WHERE week='%s'", prevWeek)
	prevWeekPlan := dbservice.GetWeeksWorkout(prevWeekSqlStatement)

	// fmt.Printf("exercise: %s\n", exerciseInfo.Name)
	// fmt.Printf("min reps: %d\n", minReps)
	// fmt.Printf("min sets: %d\n", minSets)
	// fmt.Printf("max reps: %d\n", maxReps)
	// fmt.Printf("max sets: %d\n", maxSets)
	if prevWeekPlan == nil {
		newExercisePlan = config.DeterminedWorkout{
			Name:        exerciseInfo.Name,
			Tool:        exerciseInfo.Tool,
			WeightType:  exerciseInfo.WeightType,
			Category:    exerciseInfo.Category,
			WorkoutType: exerciseInfo.WorkoutType,
			Sets:        int(minSets),
			Reps:        int(minReps),
			WorkoutId:   workoutId,
		}
	} else {
		prevWeekdayTimestamp := GetDayTimestamp(config.TDay(day), prevWeek)
		prevWeekSqlStatement := fmt.Sprintf("SELECT * FROM workouts WHERE week='%s', day='%s'", prevWeek, prevWeekdayTimestamp)
		prevWeekdayPlans := dbservice.GetWeeksWorkout(prevWeekSqlStatement)
		prevWeekExerciseInfo := GetExerciseFromPrevWeek(prevWeekdayPlans[0], workoutId)
		newReps, newSets, newWeight := prevWeekExerciseInfo.Reps, prevWeekExerciseInfo.Sets, prevWeekExerciseInfo.Weight

		if increaseVolume {
			if newReps > int(minReps) && newReps < int(maxReps) {
				newReps = CalculateIncrease(newReps, TwentyPercent)
			} else if newSets > int(minSets) && newSets < int(maxSets) {
				newSets = CalculateIncrease(newSets, TwentyPercent)
			} else {
				newWeight = CalculateIncrease(newWeight, TwentyPercent)
			}
		}

		newExercisePlan = config.DeterminedWorkout{
			Name:        exerciseInfo.Name,
			Tool:        exerciseInfo.Tool,
			WeightType:  exerciseInfo.WeightType,
			Category:    exerciseInfo.Category,
			WorkoutType: exerciseInfo.WorkoutType,
			Unit:        prevWeekExerciseInfo.Unit,
			Sets:        newSets,
			Reps:        newReps,
			Weight:      newWeight,
			WorkoutId:   workoutId,
		}
	}
	return newExercisePlan
}

func GetExerciseFromPrevWeek(plan dbservice.WorkoutTable, workoutId config.WorkoutNameIds) config.DeterminedWorkout {
	allWorkouts := append(plan.WorkoutPlan.AccessoryWorkout, plan.WorkoutPlan.MainWorkout...)
	prevWeekExercise, _ := lo.Find(allWorkouts, func(workout config.DeterminedWorkout) bool {
		return workout.WorkoutId == workoutId
	})

	return prevWeekExercise
}

func GetAccessoryWorkouts(mainWorkoutsCategories []config.WorkoutTags, increaseVolume bool, day config.TDay) []config.DeterminedWorkout {
	var allAccessories = make(config.WorkoutConfig)

	for workoutId, workoutConfig := range config.WorkoutNameConfig {
		if workoutConfig.IsAccessory == true {
			allAccessories[workoutId] = workoutConfig
		}
	}

	var categoryAccessoryIds []config.WorkoutNameIds

	for workoutId, workoutConfig := range allAccessories {
		var addCategory bool
		for _, category := range workoutConfig.Category {
			if slices.Contains(mainWorkoutsCategories, category) {
				addCategory = true
			}
		}

		if addCategory == true {
			categoryAccessoryIds = append(categoryAccessoryIds, workoutId)
		}
	}

	var accessoryWorkouts []config.DeterminedWorkout
	accessoryWorkouts = lo.Map(categoryAccessoryIds, func(workoutId config.WorkoutNameIds, index int) config.DeterminedWorkout {
		workoutDetails := ModifyExerciseDetails(workoutId, increaseVolume, day)
		return workoutDetails
	})

	return pickTwoRandomItems(accessoryWorkouts)
}

func pickTwoRandomItems(items []config.DeterminedWorkout) []config.DeterminedWorkout {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time

	n := len(items)
	if n <= 2 {
		return items // Return the original slice if it has two or fewer items
	}

	// Generate two unique random indices
	index1 := rand.Intn(n)
	index2 := rand.Intn(n - 1)
	if index2 >= index1 {
		index2++ // Increment index2 if it's greater than or equal to index1 to ensure they are unique
	}

	// Return a new slice containing the two randomly picked items
	return []config.DeterminedWorkout{items[index1], items[index2]}
}
