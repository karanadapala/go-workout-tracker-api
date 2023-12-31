package config

var WorkoutNameConfig = WorkoutConfig{
	"kb-swings": {
		Name:        "Kettlebell: Swings",
		SetRange:    "10-20",
		Tool:        "kb",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"lower-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"kb-alt-swings": {
		Name:        "Kettlebell: Alternative Hand Swings",
		SetRange:    "10-20",
		Tool:        "kb",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"lower-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"kb-clean-and-press": {
		Name:        "KettleBell: 2 Hand Clean & Press",
		SetRange:    "10-20",
		Tool:        "kb",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "3-5",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"kb-alt-clean-and-press": {
		Name:        "KettleBell: Alternative Hand Clean & Press",
		SetRange:    "10-20",
		Tool:        "kb",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "3-5",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"kb-snatches": {
		Name:        "KettleBell: 2 Hand Snatches",
		SetRange:    "10-20",
		Tool:        "kb",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "3-5",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"kb-alt-snatches": {
		Name:        "KettleBell: Alternative Hand Snatches",
		SetRange:    "10-20",
		Tool:        "kb",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "3-5",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"kb-turkish-get-up": {
		Name:        "Turkish Get Up",
		SetRange:    "10-20",
		Tool:        "kb",
		WeightType:  []WeightTags{"light"},
		Category:    []WorkoutTags{"full-body"},
		RepRange:    "2-3",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"kb-quarter-turkish-get-up": {
		Name:        "Quarter Turkish Get Up",
		SetRange:    "10-20",
		Tool:        "kb",
		WeightType:  []WeightTags{"light"},
		Category:    []WorkoutTags{"full-body"},
		RepRange:    "2-3",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"kb-half-turkish-get-up": {
		Name:        "Half Turkish Get Up",
		SetRange:    "10-20",
		Tool:        "kb",
		WeightType:  []WeightTags{"light"},
		Category:    []WorkoutTags{"full-body"},
		RepRange:    "2-3",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"mace-2h-360": {
		Name:        "Mace: 2 Hands 360s",
		SetRange:    "10-20",
		Tool:        "mace",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "10-100",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"mace-2h-10-2": {
		Name:        "Mace: 2 Hands 10-2s",
		SetRange:    "10-20",
		Tool:        "mace",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "10-100",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"mace-1h-360": {
		Name:        "Mace: 1 Hand 360",
		SetRange:    "10-20",
		Tool:        "mace",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "10-100",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"mace-1h-10-2": {
		Name:        "Mace: 1 Hand 10-2s",
		SetRange:    "10-20",
		Tool:        "mace",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "10-100",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"mace-1h-flow": {
		Name:        "Mace: 1 Hand Flow",
		SetRange:    "10-20",
		Tool:        "mace",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "10-100",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"mace-1h-inside-circle": {
		Name:        "Mace: 1 Hand Inside Circles",
		SetRange:    "10-20",
		Tool:        "mace",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"mace-1h-outside-circle": {
		Name:        "Mace: 1 Hand Outside Circles",
		SetRange:    "10-20",
		Tool:        "mace",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"club-2h-360": {
		Name:        "2 Hand Mudgar 360s",
		SetRange:    "10-20",
		Tool:        "club",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"club-2h-10-2": {
		Name:        "2 Hand Mudgar 10-2",
		SetRange:    "10-20",
		Tool:        "club",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"club-2h-inside-circle": {
		Name:        "2 Hand Mudgar Inside Circles",
		SetRange:    "10-20",
		Tool:        "club",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"club-2h-outside-circle": {
		Name:        "2 Hand Mudgar 360 Circles",
		SetRange:    "10-20",
		Tool:        "club",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"club-1h-360": {
		Name:        "1 Hand Mudgar 360s",
		SetRange:    "10-20",
		Tool:        "club",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"club-1h-10-2": {
		Name:        "1 Hand Mudgar 10-2",
		SetRange:    "10-20",
		Tool:        "club",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"club-1h-flow": {
		Name:        "1 Hand Mudgar 10-2",
		SetRange:    "10-20",
		Tool:        "club",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"club-1h-inside-circle": {
		Name:        "1 Hand Mudgar Inside Circles",
		SetRange:    "10-20",
		Tool:        "club",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"club-1h-outside-circle": {
		Name:        "1 Hand Mudgar 360 Circles",
		SetRange:    "10-20",
		Tool:        "club",
		WeightType:  []WeightTags{"heavy", "light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "emom",
		IsAccessory: false,
	},
	"squats": {
		Name:        "Bodyweight: Squats",
		SetRange:    "3-5",
		Tool:        "body-weight",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"lower-body"},
		RepRange:    "5-10",
		WorkoutType: "reps",
		IsAccessory: true,
	},
	"Dand": {
		Name:        "Bodyweight: Dand",
		SetRange:    "3-5",
		Tool:        "body-weight",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"lower-body"},
		RepRange:    "5-10",
		WorkoutType: "reps",
		IsAccessory: true,
	},
	"push-ups": {
		Name:        "Bodyweight: Push Ups",
		SetRange:    "3-5",
		Tool:        "body-weight",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "reps",
		IsAccessory: true,
	},
	"pull-ups": {
		Name:        "Bodyweight: Pull Ups",
		SetRange:    "3-5",
		Tool:        "body-weight",
		WeightType:  []WeightTags{"heavy"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "5-10",
		WorkoutType: "reps",
		IsAccessory: true,
	},
	"lateral-raises": {
		Name:        "Light Weight: Lateral Raises",
		SetRange:    "3-5",
		Tool:        "dumbbell",
		WeightType:  []WeightTags{"light"},
		Category:    []WorkoutTags{"upper-body"},
		RepRange:    "10-15",
		WorkoutType: "reps",
		IsAccessory: true,
	},
}
