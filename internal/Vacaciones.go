package internal

import "time"

type Vacaciones struct {
	Dias                uint
	Periodo             []time.Time
	Empleado            Empleado
	Restricciones       Restricciones
	VacacionesAsignadas uint8
}
