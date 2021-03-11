package db

import (
	"database/sql"

	"github.com/denizaygun/streakify/models"
)

// GetAllStreaks ...
func (db Database) GetAllStreaks() (*models.StreakList, error) {
	list := &models.StreakList{}
	rows, err := db.Conn.Query("SELECT * FROM streaks ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var streak models.Streak
		err := rows.Scan(&streak.ID, &streak.Name, &streak.Icon, &streak.Description, &streak.Count, &streak.CreatedAt, &streak.UpdatedAt)
		if err != nil {
			return list, err
		}
		list.Streaks = append(list.Streaks, streak)
	}
	return list, nil
}

// AddStreak ...
func (db Database) AddStreak(streak *models.Streak) error {
	var id int
	var createdAt string
	query := `INSERT INTO streaks (name, description, icon) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, streak.Name, streak.Description, streak.Icon).Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	streak.ID = id
	streak.CreatedAt = createdAt
	return nil
}

// GetStreakByID ...
func (db Database) GetStreakByID(streakID int) (models.Streak, error) {
	streak := models.Streak{}
	query := `SELECT * FROM streaks WHERE id = $1;`
	row := db.Conn.QueryRow(query, streakID)
	switch err := row.Scan(&streak.ID, &streak.Name, &streak.Description, &streak.Icon, &streak.Count, &streak.CreatedAt, &streak.UpdatedAt); err {
	case sql.ErrNoRows:
		return streak, ErrNoMatch
	default:
		return streak, err
	}
}

// DeleteStreak ...
func (db Database) DeleteStreak(streakID int) error {
	query := `DELETE FROM streaks WHERE id = $1;`
	_, err := db.Conn.Exec(query, streakID)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

// UpdateStreak ...
func (db Database) UpdateStreak(streakID int, streakData models.Streak) (models.Streak, error) {
	streak := models.Streak{}
	query := `UPDATE streaks SET name=$1, description=$2, count=$s3, icon=$s4, updated_at=$s5 WHERE id=$6 RETURNING id, name, description, count, icon, created_at, updated_at;`
	err := db.Conn.QueryRow(query, streakData.Name, streakData.Description, streakData.Count+1, streakData.Icon, `CURRENT_TIMESTAMP`, streakID).Scan(&streak.ID, &streak.Name, &streak.Description, &streak.Count, &streak.Icon, &streak.CreatedAt, &streak.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return streak, ErrNoMatch
		}
		return streak, err
	}
	return streak, nil
}
