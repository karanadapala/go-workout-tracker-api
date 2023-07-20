package config

var InitDBs = []string{
	`CREATE TABLE IF NOT EXISTS workouts (
		week VARCHAR(25) NOT NULL,
		day VARCHAR(25) NOT NULL,
		email VARCHAR(255) NOT NULL, 
		workout_plan TEXT,
		PRIMARY KEY (week, day, email)
	);`,
	`CREATE TABLE IF NOT EXISTS images (
		week VARCHAR(25) NOT NULL,
		day VARCHAR(25) NOT NULL,
		email VARCHAR(255) NOT NULL, 
		images blob,
		PRIMARY KEY (week, day, email)
	);`,
}
