package internal

import "time"

type Restricciones struct {
	Personas          []Empleado
	DiasImposibles    []time.Time
	Antiguedad        uint
	Peticiones        uint
	VacacionesPedidas bool
}
