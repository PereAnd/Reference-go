package core

import "github.com/jairogloz/go-l/pkg/domain"

// PlayerCreateParams es una estructura que representa los parámetros necesarios para crear un jugador.
type PlayerCreateParams struct {
	ContactInfo *domain.ContactInfo `json:"contact_info" `
	DateOfBirth *domain.Date        `json:"date_of_birth" binding:"required"`
	FirstName   string              `json:"first_name" binding:"required"`
	LastName    string              `json:"last_name" binding:"required"`
	TeamInfo    *domain.TeamInfo    `json:"team_info" binding:"required"`
}
