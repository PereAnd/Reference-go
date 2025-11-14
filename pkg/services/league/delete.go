package league

import (
	"context"
	"errors"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Delete elimina una liga por su ID.
func (s *Service) Delete(ctx context.Context, id string) (err error) {
    if id == "" {
        return errors.New("id is required")
    }
    err = s.Repo.Delete(ctx, id)
    if err != nil {
        return domain.ManageError(err, "Error deleting league")
    }
    return nil
}
