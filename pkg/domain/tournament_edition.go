package domain

import "time"

const (
	TournamentTypeLeague  = "league"   // liga, ej. La Liga en España o la Premier League en Inglaterra
	TournamentTypeCup     = "knockout" // eliminatorias, ej. Copa del Rey en España o FA Cup en Inglaterra
	TournamentTypePlayoff = "playoff"  // playoff, Etapa regular seguido de playoff, ej. Mundial de la Fifa
)

// TournamentEdition representa una edición de torneo deportivo en nuestro sistema.
// Una edición es una instancia específica de un torneo, por ejemplo, la Liga
// Golang League puede tener un Torneo llamado Kids Tournament, y este Torneo
// puede tener una TournamentEdition llamada 2024 Edition, luego el próximo año el mismo Torneo
// puede tener una nueva TournamentEdition llamada 2025 Edition.
//
// Una TournamentEdition puede tener múltiples equipos que participan en ella, pueden variar entre
// ediciones.
type TournamentEdition struct {
	CreatedAt     *time.Time     `json:"created_at" bson:"created_at"`
	Description   string         `json:"description" bson:"description"`
	EndDate       *time.Time     `json:"end_date" bson:"end_date"`
	GameDays      []time.Weekday `json:"game_days" bson:"game_days"`
	GameTimeRange *TimeRange     `json:"game_time_range" bson:"game_time_range"`
	ID            string         `json:"id" bson:"_id,omitempty"`
	Name          string         `json:"name" bson:"name"`
	StartDate     *time.Time     `json:"start_date" bson:"start_date"`
	Type          string         `json:"type" bson:"type"`
	UpdatedAt     *time.Time     `json:"updated_at" bson:"updated_at"`
}
