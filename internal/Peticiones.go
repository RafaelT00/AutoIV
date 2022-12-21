package internal

import "time"

type Season uint8

const (
	Invierno Season = 0
	Primavera 		= 1
	Verano 			= 2
	Otonio			= 3
)

type Peticiones struct{
	DiasQueQuiero []Time.time
	Temporada Season
}