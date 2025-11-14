package domain

import "time"

// League representa una liga deportiva en nuestro sistema.
// Una liga puede tener uno o más Torneos, y un Torneo puede tener una o más
// Ediciones, por ejemplo, la Liga Golang League puede tener un Torneo llamado
// Kids Tournament, y este Torneo puede tener una Edición llamada 2024 Edition,
// luego el próximo año el mismo Torneo puede tener una nueva Edición llamada 2025 Edition.
type League struct {
	CreatedAt   *time.Time  `json:"created_at" bson:"created_at"`
	Description string      `json:"description" bson:"description"`
	ID          interface{} `json:"id" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	UpdatedAt   *time.Time  `json:"updated_at" bson:"updated_at"`
}
