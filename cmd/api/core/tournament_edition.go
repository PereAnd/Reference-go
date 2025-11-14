package core

import (
	"time"

	"github.com/jairogloz/go-l/pkg/domain"
)

// TournamentEditionParams es una estructura que representa los parámetros necesarios para editar un torneo.
type TournamentEditionParams struct {
	Description   string            `json:"description" binding:"required"`
	EndDate       *time.Time        `json:"end_date"` // Debe ser una fecha válida en el futuro, después de StartDate.
	GameDays      []time.Weekday    `json:"game_days" binding:"required"`
	GameTimeRange *domain.TimeRange `json:"game_time_range" bson:"game_time_range"`
	Name          string            `json:"name" binding:"required"`
	StartDate     *time.Time        `json:"start_date" bson:"start_date"`
	Type          string            `json:"type" bson:"type"`
	URL           string            `json:"url" binding:"required"`
}
