package internal

import "time"

type Tarea uint8

const (
	Facil   Tarea = 0
	Medio         = 1
	Dificil       = 2
	Urgente       = 3
)

type Asignar struct {
	tareas              []Tarea
	peticiones          Peticiones
	empleado            Empleado
	vacacionesAsignadas []time.Time
}
