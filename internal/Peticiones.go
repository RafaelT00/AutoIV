package internal

type Season uint8
type Semana uint8

const (
	Invierno  Season = 0
	Primavera        = 1
	Verano           = 2
	Otonio           = 3
)

const (
	Lunes Semana = iota //0
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
	Domingo //6
)

type Peticiones struct {
	diasQueQuiero uint
	temporada     Season
	dia           Semana
}

// s.String() devuelve el dia de la semana en string en vez de en numero
func (s Semana) String() string {
	return []string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}[s]
}
