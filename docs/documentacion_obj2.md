# Documentación sobre la implementación
## Análisis del dominio
Analizamos la HU RafaelT00#7. Es necesario entender como se van a relacionar las distintas partes ahí escritas. Primero para dividir la historia de usuario en trozos más pequeños para entenderlos y luego relacionarlos, obtengo los encargados de asignar las tareas, en los que me surge un problema de como contextualizarlos RafaelT00#19, en el cual creo una solución que se puede adaptar [ecbb97e](https://github.com/ignaciotitos/AutoIV/commit/ecbb97ebcae642762dc5848340d475ad89300afd). Más tarde con la revisión el compañero sugirió cambios en RafaelT00#19 por lo que cambié la solución propuesta por la siguiente cerrando el issue [69f8b4d](https://github.com/ignaciotitos/AutoIV/commit/69f8b4d034cfc44d002fa6046b12d9fea5e205dc).

El siguiente problema en RafaelT00#7 que me encontré son las tareas, en las que tengo que comprender como cuantificar las tareas importantes, cual es su dificultad, cuantas tareas tiene un empleado. Planteado en RafaelT00#20. En este problema planteo una solución que la propongo con este commit [17f2199](https://github.com/ignaciotitos/AutoIV/commit/17f2199aa10ca01182bf7c21eb058762507fe183)

El problema siguiente en RafaelT00#7 fue de como se iban a gestionar las peticiones por parte de los encargados, cuando son las épocas de mayores peticiones y como afectan a las asignaciones. Este problema está referenciado en este issue RafaelT00#21, en el que añadí una solución en los commmits 
[17f2199](https://github.com/ignaciotitos/AutoIV/commit/17f2199aa10ca01182bf7c21eb058762507fe183),
[3a5c699](https://github.com/ignaciotitos/AutoIV/commit/3a5c69914382e0ff4853c232d184ce7706a0cad7)

Resuelto el problema del manejo de las peticiones, aparece otro problema planteado en RafaelT00#24 el cual como saber si las peticiones son X días y lo pides un viernes por ejemplo, que no te cuente las vacaciones ni los días no laborables, por lo que se ha hecho uso de un enumerable con los días de la semana para que cuando se lleven los días, y sepas que día lo pidió y los días que pidió de vacaciones, puedas hacer una asignación real. [7a24cff](https://github.com/ignaciotitos/AutoIV/commit/7a24cff122951ebfdc01d8bd84dbcda3713af3af)

El siguiente problema a resolver fue de cuanto va a ser la duración de cada proyecto RafaelT00#25, para poder asignar días reales a un empleado. Para esto se hace uso de la durabilidad del proyecto en días (sprint/s), para poder hacer una asignación real. Solución propusta en [87b2b13](https://github.com/ignaciotitos/AutoIV/commit/87b2b13db1db49609e085fe3556374a2cf9dc3ab)

## Objeto valor o entidad
Para cada clase ha sido necesario saber cómo va a funcionar cada clase, para saber si van a ser inmutables o mutables. Para esto ha sido necesario entender el problema para estructurar las clases. He propuesto una solución en el que las entidades sean Peticiones y Proyecto y el objeto valor sea Empleado.
