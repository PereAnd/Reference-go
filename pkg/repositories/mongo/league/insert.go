package league

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Insert persiste un nuevo documento de liga en la colección de MongoDB.
// Genera un nuevo ObjectID para la liga si no está establecido, luego inserta
// el documento. Si ocurre un error de clave duplicada, retorna un error específico
// del dominio ErrDuplicateKey. Otros errores son envueltos y retornados.
func (r *Repository) Insert(ctx context.Context, league *domain.League) (err error) {
	league.ID = primitive.NewObjectID()

	_, err = r.Collection.InsertOne(ctx, league)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Println("Duplicate key error")
			return fmt.Errorf("%w: error inserting league: %s",
				domain.ErrDuplicateKey, err.Error())
		}
		log.Println(err.Error())
		return fmt.Errorf("error inserting league: %w", err)
	}

	return nil
}
