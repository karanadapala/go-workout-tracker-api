package config

var BaseTemplate = BaseWeeklyTemplate{
	"monday": BaseWorkDay{
		MainWorkout: []WorkoutNameIds{
			"kb-alt-swings",
			"kb-alt-clean-and-press",
		},
	},
	"tuesday": BaseWorkDay{
		MainWorkout: []WorkoutNameIds{
			"club-2h-360",
			"club-2h-inside-circle",
		},
	},
	"wednesday": BaseWorkDay{
		MainWorkout: []WorkoutNameIds{
			"mace-1h-10-2",
			"mace-1h-360",
		},
	},
	"thursday": BaseWorkDay{
		MainWorkout: []WorkoutNameIds{
			"kb-half-turkish-get-up",
			"kb-alt-snatches",
		},
	},
	"friday": BaseWorkDay{
		MainWorkout: []WorkoutNameIds{
			"club-2h-10-2",
			"club-2h-outside-circle",
		},
	},
	"saturday": BaseWorkDay{
		MainWorkout: []WorkoutNameIds{
			"mace-1h-10-2",
			"mace-1h-360",
		},
	},
	"sunday": BaseRestDay{
		IsRestDay: true,
	},
}
