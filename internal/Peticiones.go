package internal

import "time"

type Season uint8

const (
	Invierno Season = 1
	Primavera 		= 2
	Verano 			= 3
	Otonio			= 4
)

type Peticiones struct{
	DiasQueQuiero []Time.time
	Temporada Season
}