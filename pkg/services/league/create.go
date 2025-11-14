package league

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Create crea una nueva liga en el sistema.
// Establece el timestamp CreatedAt y delega la persistencia al repositorio.
// Si ocurre un error de clave duplicada, retorna un AppError específico del dominio.
// Cualquier otro error de persistencia es envuelto y retornado.
func (s *Service) Create(ctx context.Context, league *domain.League) (err error) {
	now := time.Now().UTC()
	league.CreatedAt = &now

	err = s.Repo.Insert(ctx, league)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateKey) {
			log.Println("Duplicate key error")
			appErr := domain.AppError{
				Code: domain.ErrCodeDuplicateKey,
				Msg:  "error creating league: duplicate key error",
			}
			return appErr
		}
		log.Println(err.Error())
		return fmt.Errorf("error creating league: %w", err)
	}

	return
}
