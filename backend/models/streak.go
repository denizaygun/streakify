package models

import (
	"fmt"
	"net/http"
)

// Streak structure
type Streak struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Count       int    `json:"count"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// StreakList structure
type StreakList struct {
	Streaks []Streak `json:"streaks"`
}

// Bind ...
func (i *Streak) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

// Render a StreakList...
func (*StreakList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render a Streak...
func (*Streak) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
