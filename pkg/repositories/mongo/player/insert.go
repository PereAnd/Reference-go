package player

import (
	"context"
	"fmt"
	"log"

	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Insert persiste un nuevo documento de jugador en la colección de MongoDB.
// Genera un nuevo ObjectID para el jugador si no está establecido, luego inserta
// el documento. Si ocurre un error de clave duplicada, retorna un error específico
// del dominio ErrDuplicateKey. Otros errores son envueltos y retornados.
func (r *Repository) Insert(ctx context.Context, player *domain.Player) (err error) {

	player.ID = primitive.NewObjectID()

	_, err = r.Collection.InsertOne(ctx, player)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Println("Duplicate key error")
			return fmt.Errorf("%w: error inserting player: %s",
				domain.ErrDuplicateKey, err.Error())
		}
		log.Println(err.Error())
		return fmt.Errorf("error inserting player: %w", err)
	}

	return nil
}
