package AutoIV

import (
	fmt
	time
)

type Restricciones struct{
	NombreProyecto string
	Personas []string
	DiasImposibles []time.Time
}