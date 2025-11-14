package team

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Get obtiene un equipo de la base de datos, o retorna un error si algo sale mal.
func (s Service) Get(ctx context.Context, id string) (team *domain.Team, err error) {
	team, err = s.Repo.Get(ctx, id)

	if err != nil {
		return nil, domain.ManageError(err, "unexpected error getting team")
	}

	return team, nil
}
