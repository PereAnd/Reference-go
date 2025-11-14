package domain

import (
	"time"
)

// Team representa un equipo deportivo en nuestro sistema.
type Team struct {
	CreatedAt *time.Time  `json:"created_at" bson:"created_at"`
	ID        interface{} `json:"-" bson:"_id,omitempty"`
	Name      string      `json:"name" bson:"name"`
	UpdatedAt *time.Time  `json:"updated_at" bson:"updated_at"`
}

// TeamInfo representa la información del equipo de un jugador, no la
// información completa del equipo.
type TeamInfo struct {
	TeamID       string `json:"team_id" bson:"team_id"`
	JerseyNumber int    `json:"jersey_number" bson:"jersey_number"`
}
