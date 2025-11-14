package team

import (
	"context"
	"time"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Create crea un nuevo equipo basado en los datos proporcionados y lo almacena en la base de datos.
func (s Service) Create(ctx context.Context, team *domain.Team) (err error) {
	now := time.Now().UTC()
	team.CreatedAt = &now

	err = s.Repo.Insert(ctx, team)
	if err != nil {
		return domain.ManageError(err, "Error creating team")
	}
	return nil
}
