package internal

type Tarea uint8

const (
	Facil   Tarea = 0
	Medio         = 1
	Dificil       = 2
	Urgente       = 3
)

type Manejar struct {
	Tareas     []Tarea
	Peticiones Peticiones
}
