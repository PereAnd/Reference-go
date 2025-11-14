package core

// TournamentCreateParams es una estructura que representa los parámetros necesarios para crear un torneo.
type TournamentCreateParams struct {
	Description string `json:"description" binding:"required"`
	Name        string `json:"name" binding:"required"`
	URL         string `json:"url" binding:"required"`
}
