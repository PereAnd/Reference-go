package team

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get obtiene un equipo de la colección de MongoDB por su ID.
// Primero convierte el ID de cadena proporcionado a un ObjectID de MongoDB.
// Si la conversión falla (ej., si el ID no es una cadena hexadecimal válida), retorna un error.
// Luego intenta encontrar un documento en la colección de MongoDB con el ObjectID convertido.
// Si encuentra un documento, lo decodifica en un objeto domain.Team.
// Si no encuentra un documento, retorna un error específico del dominio de no encontrado.
// Si hay un error de timeout al acceder a la colección de MongoDB, retorna un error específico del dominio de timeout.
// Para cualquier otro error, retorna el error tal cual.
// Si encuentra y decodifica exitosamente un documento, establece el ID del objeto domain.Team
// a la representación hexadecimal del ObjectID y retorna el objeto domain.Team.
func (r Repository) Get(ctx context.Context, id string) (team *domain.Team, err error) {
	teamID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return team, fmt.Errorf("%w: error getting team: %s", domain.ErrIncorrectID, err)
	}

	err = r.Collection.FindOne(ctx, bson.M{"_id": teamID}).Decode(&team)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}

		var commandErr *mongo.CommandError
		if errors.As(err, &commandErr) && commandErr.HasErrorLabel("NetworkTimeout") {
			return nil, domain.ErrTimeout
		}
		return nil, err
	}

	team.ID = teamID.Hex()
	return team, nil
}
