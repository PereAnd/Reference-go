package domain

// Date representa una fecha con día, mes y año.
type Date struct {
	Day   int `json:"day" bson:"day"`
	Month int `json:"month" bson:"month"`
	Year  int `json:"year" bson:"year"`
}

// TimeRange representa un rango de tiempo con horas de inicio y fin. Útil
// para especificar el rango de tiempo de los juegos.
type TimeRange struct {
	Start TimeOfDay `json:"start" bson:"start"`
	End   TimeOfDay `json:"end" bson:"end"`
}

// TimeOfDay representa una hora del día con hora y minuto.
type TimeOfDay struct {
	Hour   int `json:"hour" bson:"hour"`
	Minute int `json:"minute" bson:"minute"`
}
