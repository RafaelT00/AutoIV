package src

import (
	fmt
	Vacaciones
)

type Empleado struct{
	Nombre string
	Ficha string
	DiasLibres []Vacaciones
	Peticiones uint
	Antiguedad uint
	VacacionesElegidas bool
}
