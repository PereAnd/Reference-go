package player

import (
	"context"
	"fmt"

	"github.com/jairogloz/go-l/pkg/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Asegura que Repository implementa ports.PlayerRepository en tiempo de compilación.
var _ ports.PlayerRepository = &Repository{}

// Repository implementa la interfaz PlayerRepository usando MongoDB como mecanismo de persistencia.
// Proporciona métodos para insertar, recuperar y eliminar documentos de jugadores en MongoDB.
type Repository struct {
	// Client es el cliente de MongoDB usado para operaciones de base de datos.
	Client *mongo.Client

	// Collection es la colección de MongoDB donde se almacenan los documentos de jugadores.
	Collection *mongo.Collection
}

// CreateIndexes crea índices secundarios para la colección de jugadores para optimizar consultas.
// Actualmente crea un índice en team_info.team_id para búsquedas eficientes.
// Este método debe ser llamado durante la inicialización de la aplicación.
func (r *Repository) CreateIndexes() error {
	// Crea el índice team_id
	// Este índice se usa para buscar jugadores por team_id
	// El índice no es único porque un equipo puede tener múltiples jugadores
	_, err := r.Collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "team_info.team_id", Value: 1}},
		Options: options.Index().SetUnique(false),
	})
	if err != nil {
		return fmt.Errorf("error creating team_id index: %w", err)
	}

	return nil
}
