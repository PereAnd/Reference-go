package tournament

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Delete elimina un torneo por su ID de la base de datos.
func (r *Repository) Delete(ctx context.Context, id string) (err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrIncorrectID
	}

	deleteResult, err := r.Collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return fmt.Errorf("error deleting tournament: %s", err.Error())
	}

	if deleteResult.DeletedCount == 0 {
		return domain.ErrNotFound
	}

	return nil
}
