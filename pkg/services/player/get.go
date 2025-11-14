package player

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Get obtiene un jugador por su ID desde el repositorio.
// Retorna un objeto domain.Player y un error.
// Si el ID proporcionado está vacío, retorna un error.
// Si el jugador no se encuentra en el repositorio, retorna un error específico del dominio de no encontrado.
// Si hay un error de timeout al acceder al repositorio, retorna un error específico del dominio de timeout.
// Para cualquier otro error, registra el error y retorna un error genérico.
func (s *Service) Get(ctx context.Context, id string) (player *domain.Player, err error) {

	if id == "" {
		return nil, errors.New("id is required")
	}

	player, err = s.Repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, domain.NewAppError(
				domain.ErrCodeNotFound,
				fmt.Sprintf("player with id '%s' not found", id))
		}
		if errors.Is(err, domain.ErrTimeout) {
			return nil, domain.NewAppError(
				domain.ErrCodeTimeout,
				"timeout error, try again later")
		}
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting player: %w", err)
	}

	return player, nil
}
