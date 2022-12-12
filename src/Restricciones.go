package src

import (
	fmt
	time
  Empleado
)

type Restricciones struct{
	NombreProyecto string
	Personas []Empleado
	DiasImposibles []time.Time
}
