package config

type TDay string

const (
	Monday    TDay = "monday"
	Tuesday   TDay = "tuesday"
	Wednesday TDay = "wednesday"
	Thursday  TDay = "thursday"
	Friday    TDay = "friday"
	Saturday  TDay = "saturday"
	Sunday    TDay = "sunday"
)

type Tools string

const (
	Kbs  Tools = "kbs"
	Mace Tools = "mace"
	Club Tools = "club"
)

type WorkoutType string

const (
	Emom    WorkoutType = "emom"
	ForTime WorkoutType = "forTime"
	Reps    WorkoutType = "reps"
)

type Workout struct {
	Name   WorkoutNameIds
	Reps   int
	Sets   int
	Weight int
	Type   WorkoutType
}

type WorkoutNameIds string

const (
	KbSwings              WorkoutNameIds = "kb-swings"
	KbAltSwings           WorkoutNameIds = "kb-alt-swings"
	KbCleanAndPress       WorkoutNameIds = "kb-clean-and-press"
	KbAltCleanAndPress    WorkoutNameIds = "kb-alt-clean-and-press"
	KbSnatches            WorkoutNameIds = "kb-snatches"
	KbAltSnatches         WorkoutNameIds = "kb-alt-snatches"
	KbTurkishGetUp        WorkoutNameIds = "kb-turkish-get-up"
	KbQuarterTurkishGetUp WorkoutNameIds = "kb-quarter-turkish-get-up"
	KbHalfTurkishGetUp    WorkoutNameIds = "kb-half-turkish-get-up"
	Mace2H360             WorkoutNameIds = "mace-2h-360"
	Mace2H102             WorkoutNameIds = "mace-2h-10-2"
	Mace1H360             WorkoutNameIds = "mace-1h-360"
	Mace1H102             WorkoutNameIds = "mace-1h-10-2"
	Mace1HFlow            WorkoutNameIds = "mace-1h-flow"
	Mace1HInsideCircle    WorkoutNameIds = "mace-1h-inside-circle"
	Mace1HOutsideCircle   WorkoutNameIds = "mace-1h-outside-circle"
	Club2H360             WorkoutNameIds = "club-2h-360"
	Club2H102             WorkoutNameIds = "club-2h-10-2"
	Club2HInsideCircle    WorkoutNameIds = "club-2h-inside-circle"
	Club2HOutsideCircle   WorkoutNameIds = "club-2h-outside-circle"
	Club1H360             WorkoutNameIds = "club-1h-360"
	Club1H102             WorkoutNameIds = "club-1h-10-2"
	Club1HFlow            WorkoutNameIds = "club-1h-flow"
	Club1HInsideCircle    WorkoutNameIds = "club-1h-inside-circle"
	Club1HOutsideCircle   WorkoutNameIds = "club-1h-outside-circle"
	Squats                WorkoutNameIds = "squats"
	Dand                  WorkoutNameIds = "Dand"
	PushUps               WorkoutNameIds = "push-ups"
	PullUps               WorkoutNameIds = "pull-ups"
	LateralRaises         WorkoutNameIds = "lateral-raises"
)

type WorkoutConfigValue struct {
	Name        string
	Tool        ToolTags
	Category    []WorkoutTags
	WeightType  []WeightTags
	IsAccessory bool
	RepRange    RepRange
	WorkoutType WorkoutType
	SetRange    SetRange
}

type WorkoutConfig map[WorkoutNameIds]WorkoutConfigValue

type ToolTags string

const (
	Kb         ToolTags = "kb"
	Clubbell   ToolTags = "club"
	Macebell   ToolTags = "mace"
	Dumbbell   ToolTags = "dumbbell"
	BodyWeight ToolTags = "body-weight"
)

type WeightTags string

const (
	Light WeightTags = "light"
	Heavy WeightTags = "heavy"
)

type RepRange string
type SetRange string

type WorkoutTags string

const (
	UpperBody WorkoutTags = "upper-body"
	LowerBody WorkoutTags = "lower-body"
	FullBody  WorkoutTags = "full-body"
)

type WarmUpId string
type CoolDownId string

type WarmUpValue struct {
	Label    string
	Reps     int
	Time     int // 1 minute
	Category WorkoutTags
}

type WarmUp map[WarmUpId]WarmUpValue

type CoolDownValue struct {
	Label    string
	Reps     int
	Time     int // 1 minute
	Category WorkoutTags
}

type CoolDown map[CoolDownId]CoolDownValue

type DeterminedWorkout struct {
	Name        string
	Sets        int
	Tool        ToolTags
	WeightType  []WeightTags
	Weight      int
	Unit        string
	Category    []WorkoutTags
	Reps        int
	WorkoutType WorkoutType
	WorkoutId   WorkoutNameIds
}

type DailyPlan struct {
	WarmUp           []WarmUpValue
	MainWorkout      []DeterminedWorkout
	AccessoryWorkout []DeterminedWorkout
	CoolDown         []CoolDownValue
}

type WeeklyPlan map[TDay]DailyPlan

type BaseWorkDay struct {
	MainWorkout []WorkoutNameIds
}

type BaseRestDay struct {
	IsRestDay bool
}

// Define the interface for BaseWeeklyTemplate
type BaseWeeklyTemplate map[TDay]BaseWorkDay
