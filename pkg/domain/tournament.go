package domain

import "time"

// Tournament representa un torneo deportivo en una Liga en nuestro sistema.
// Un torneo puede tener una o más Ediciones, por ejemplo, la Liga Golang League
// puede tener un Torneo llamado Kids Tournament, y este Torneo puede tener una Edición
// llamada 2024 Edition, luego el próximo año el mismo Torneo puede tener una nueva Edición
// llamada 2025 Edition.
type Tournament struct {
	CreatedAt   *time.Time  `json:"created_at" bson:"created_at"`
	Description string      `json:"description" bson:"description"`
	ID          interface{} `json:"id" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	URL         string      `json:"url" bson:"url"`
	UpdatedAt   *time.Time  `json:"updated_at" bson:"updated_at"`
}
