package core

// LeagueCreateParams es una estructura que representa los parámetros necesarios para crear una liga.
type LeagueCreateParams struct {
	Description string `json:"description" binding:"required"`
	Name        string `json:"name" binding:"required"`
}
