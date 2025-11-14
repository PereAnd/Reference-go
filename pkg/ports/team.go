package ports

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// TeamService define la interfaz del servicio de aplicación para operaciones de equipos.
// Este es un puerto conductor (entrada) que define las operaciones de lógica de negocio
// disponibles para equipos. Las implementaciones deben contener reglas de negocio y
// orquestar operaciones del dominio.
type TeamService interface {
	// Create crea un nuevo equipo en el sistema.
	// Valida los datos del equipo y aplica reglas de negocio antes de persistir.
	// Retorna un error si el equipo no puede ser creado (ej., clave duplicada, fallo de validación).
	Create(ctx context.Context, team *domain.Team) (err error)

	// Get obtiene un equipo por su ID.
	// Retorna el equipo si se encuentra, o un error si no se encuentra o si la operación falla.
	Get(ctx context.Context, id string) (team *domain.Team, err error)

	// Delete elimina un equipo del sistema por su ID.
	// Retorna un error si el equipo no se encuentra o si la eliminación falla.
	Delete(ctx context.Context, id string) (err error)
}

// TeamRepository define la interfaz del repositorio para operaciones de persistencia de equipos.
// Este es un puerto conducido (salida) que abstrae las operaciones de acceso a datos.
// Las implementaciones manejan el mecanismo de persistencia real (ej., MongoDB, PostgreSQL).
type TeamRepository interface {
	// Insert persiste un nuevo equipo en el almacén de datos.
	// El ID del equipo debe ser generado si no se proporciona.
	// Retorna un error si la inserción falla (ej., clave duplicada, error de conexión).
	Insert(ctx context.Context, team *domain.Team) (err error)

	// Get obtiene un equipo del almacén de datos por su ID.
	// Retorna el equipo si se encuentra, o un error si no se encuentra o si la operación falla.
	Get(ctx context.Context, id string) (team *domain.Team, err error)

	// Delete elimina un equipo del almacén de datos por su ID.
	// Retorna un error si el equipo no se encuentra o si la eliminación falla.
	Delete(ctx context.Context, id string) (err error)
}
