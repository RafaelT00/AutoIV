package internal

import "time"

type Restricciones struct {
	DiasImposibles    []time.Time
	Antiguedad        uint
	Peticiones        uint
	VacacionesPedidas bool
}
