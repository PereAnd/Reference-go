package domain

// ContactInfo representa la información de contacto de un jugador.
type ContactInfo struct {
	Email string `json:"email" bson:"email" binding:"email"`
	Phone string `json:"phone" bson:"phone"`
}
