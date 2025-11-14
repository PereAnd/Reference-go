package ports

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// TournamentService define la interfaz del servicio de aplicación para operaciones de torneos.
// Este es un puerto conductor (entrada) que define las operaciones de lógica de negocio
// disponibles para torneos. Las implementaciones deben contener reglas de negocio y
// orquestar operaciones del dominio.
type TournamentService interface {
	// Create crea un nuevo torneo en el sistema.
	// Valida los datos del torneo y aplica reglas de negocio antes de persistir.
	// Retorna un error si el torneo no puede ser creado (ej., clave duplicada, fallo de validación).
	Create(ctx context.Context, tournament *domain.Tournament) (err error)

	// Delete elimina un torneo del sistema por su ID.
	// Retorna un error si el torneo no se encuentra o si la eliminación falla.
	Delete(ctx context.Context, id string) (err error)
}

// TournamentRepository define la interfaz del repositorio para operaciones de persistencia de torneos.
// Este es un puerto conducido (salida) que abstrae las operaciones de acceso a datos.
// Las implementaciones manejan el mecanismo de persistencia real (ej., MongoDB, PostgreSQL).
type TournamentRepository interface {
	// Insert persiste un nuevo torneo en el almacén de datos.
	// El ID del torneo debe ser generado si no se proporciona.
	// Retorna un error si la inserción falla (ej., clave duplicada, error de conexión).
	Insert(ctx context.Context, tournament *domain.Tournament) (err error)

	// Delete elimina un torneo del almacén de datos por su ID.
	// Retorna un error si el torneo no se encuentra o si la eliminación falla.
	Delete(ctx context.Context, id string) (err error)
}
