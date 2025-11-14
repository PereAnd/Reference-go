// Package league proporciona la implementación del repositorio MongoDB para persistencia de ligas.
// Este paquete implementa la interfaz LeagueRepository, sirviendo como un adaptador conducido
// (salida) en la arquitectura hexagonal. Maneja todas las operaciones específicas de MongoDB
// para entidades de ligas, traduciendo entre modelos de dominio y documentos de base de datos.
package league

import (
	"context"
	"fmt"

	"github.com/jairogloz/go-l/pkg/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Asegura que Repository implementa ports.LeagueRepository en tiempo de compilación.
var _ ports.LeagueRepository = &Repository{}

// Repository implementa la interfaz LeagueRepository usando MongoDB como mecanismo de persistencia.
// Proporciona métodos para insertar, recuperar y eliminar documentos de ligas en MongoDB.
type Repository struct {
	// Client es el cliente de MongoDB usado para operaciones de base de datos.
	Client *mongo.Client

	// Collection es la colección de MongoDB donde se almacenan los documentos de ligas.
	Collection *mongo.Collection
}

// CreateIndexes crea índices secundarios para la colección de ligas para optimizar consultas.
// Actualmente crea un índice en league_info.league_id para búsquedas eficientes.
// Este método debe ser llamado durante la inicialización de la aplicación.
func (r *Repository) CreateIndexes() error {
	// Crea el índice league_id
	// Este índice se usa para buscar ligas por league_id
	// El índice no es único porque una liga puede tener múltiples registros
	_, err := r.Collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "league_info.league_id", Value: 1}},
		Options: options.Index().SetUnique(false),
	})
	if err != nil {
		return fmt.Errorf("error creating league_id index: %w", err)
	}

	// Agregar índices adicionales si es necesario

	return nil
}
