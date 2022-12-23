package internal

type Season string
type Semana string

const (
	Invierno  Season = "Invierno"
	Primavera        = "Primavera"
	Verano           = "Verano"
	Otonio           = "Otonio"
)

const (
	Lunes 	  Semana = "Lunes"
	Martes		 	 = "Martes"
	Miercoles	 	 = "Miercoles"
	Jueves 		 	 = "Jueves"
	Viernes 	 	 = "Viernes"
	Sabado 		 	 = "Sabado"
	Domingo 	 	 = "Domingo"
)

type Peticiones struct {
	diasQueQuiero uint
	temporada     Season
	dia           Semana
}
