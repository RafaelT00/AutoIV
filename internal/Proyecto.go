package internal

type Tarea uint8

const (
	Facil   Tarea = 0
	Medio         = 1
	Dificil       = 2
	Urgente       = 3
)

type Proyecto struct {
	tareas           []Tarea
	empleadosEnTarea []Empleado
	durabilidad      uint8
}
