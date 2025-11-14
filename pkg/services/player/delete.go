package player

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Delete elimina un jugador por su ID.
func (s *Service) Delete(ctx context.Context, id string) (err error) {
	err = s.Repo.Delete(ctx, id)
	if err != nil {
		return domain.ManageError(err, "Error deleting player")
	}
	return nil
}
