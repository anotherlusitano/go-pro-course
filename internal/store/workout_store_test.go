package store

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	var (
		host     = "localhost"
		user     = "postgres"
		password = "postgres"
		dbname   = "postgres"
		port     = "5433"
		sslmode  = "disable"
	)

	dataSource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		t.Fatalf("Opening test db error: %v", err)
	}

	err = Migrate(db, "../../migrations/")
	if err != nil {
		t.Fatalf("Migrating test db error: %v", err)
	}

	_, err = db.Exec(`TRUNCATE workouts, workout_entries CASCADE`)
	if err != nil {
		t.Fatalf("Truncating tables error: %v", err)
	}

	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)

	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "Valid workout",
			workout: &Workout{
				Title:           "Push Day",
				Description:     "Workout for chest and triceps",
				DurationMinutes: 60,
				CaloriesBurned:  500,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Bench Press",
						Sets:         4,
						Reps:         IntPtr(10),
						Weight:       FloatPtr(100.0),
						Notes:        "Felt strong today",
						OrderIndex:   1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid workout",
			workout: &Workout{
				Title:           "Full Body",
				Description:     "Complete workout for all muscle groups",
				DurationMinutes: 10,
				CaloriesBurned:  8000,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Squats",
						Sets:         2,
						Reps:         IntPtr(80),
						Weight:       FloatPtr(500.0),
						Notes:        "Felt strong today",
						OrderIndex:   1,
					},
					{
						ExerciseName:    "Deadlifts",
						Sets:            2,
						Reps:            IntPtr(40),
						DurationSeconds: IntPtr(90),
						Weight:          FloatPtr(600.0),
						Notes:           "Felt strong today",
						OrderIndex:      2,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdWorkout, err := store.CreateWorkout(tt.workout)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.workout.Title, createdWorkout.Title)
			assert.Equal(t, tt.workout.Description, createdWorkout.Description)
			assert.Equal(t, tt.workout.DurationMinutes, createdWorkout.DurationMinutes)

			retrived, err := store.GetWorkoutByID(int64(createdWorkout.ID))
			require.NoError(t, err)

			assert.Equal(t, createdWorkout.ID, retrived.ID)
			assert.Equal(t, len(tt.workout.Entries), len(retrived.Entries))

			for i := range retrived.Entries {
				assert.Equal(t, tt.workout.Entries[i].ExerciseName, retrived.Entries[i].ExerciseName)
				assert.Equal(t, tt.workout.Entries[i].Sets, retrived.Entries[i].Sets)
				assert.Equal(t, tt.workout.Entries[i].OrderIndex, retrived.Entries[i].OrderIndex)
			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}

func FloatPtr(f float64) *float64 {
	return &f
}
