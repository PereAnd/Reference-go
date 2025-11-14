package ports

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

type PlayerService interface {
	Create(ctx context.Context, player *domain.Player) (err error)
	Get(ctx context.Context, id string) (player *domain.Player, err error)
	Delete(ctx context.Context, id string) (err error)
}

type PlayerRepository interface {
	Insert(ctx context.Context, player *domain.Player) (err error)
	Get(ctx context.Context, id string) (player *domain.Player, err error)
	GetPlayersByTeamID(ctx context.Context, id string) (players []*domain.Player, err error)
	Delete(ctx context.Context, id string) (err error)
}
