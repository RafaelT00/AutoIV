# Documentación sobre la implementación
## Análisis del dominio
Analizamos la HU RafaelT00#7. Es necesario entender como se van a relacionar las distintas partes ahí escritas. Primero para dividir la historia de usuario en trozos más pequeños para entenderlos y luego relacionarlos, obtengo los encargados de asignar las tareas, en los que me surge un problema de como contextualizarlos RafaelT00#19, en el cual creo una solución que se puede adaptar [ecbb97e](https://github.com/ignaciotitos/AutoIV/commit/ecbb97ebcae642762dc5848340d475ad89300afd). Más tarde con la revisión el compañero sugirió cambios en RafaelT00#19 por lo que cambié la solución propuesta por la siguiente cerrando el issue [69f8b4d](https://github.com/ignaciotitos/AutoIV/commit/69f8b4d034cfc44d002fa6046b12d9fea5e205dc).

Después el siguiente problema que me encontré son las tareas, en las que tengo que comprender como cuantificar las tareas importantes, cual es su dificultad, cuantas tareas tiene un empleado. Planteado en RafaelT00#20. En este problema planteo una solución que la propongo con este commit [17f2199](https://github.com/ignaciotitos/AutoIV/commit/17f2199aa10ca01182bf7c21eb058762507fe183)

Más tarde surgió el problema de como se iban a gestionar las peticiones por parte de los encargados, cuando son las épocas de mayores peticiones y como afectan a las asignaciones. Este problema está referenciado en este issue RafaelT00#21, en el que añadí una solución en los commmits 
[17f2199](https://github.com/ignaciotitos/AutoIV/commit/17f2199aa10ca01182bf7c21eb058762507fe183),
[3a5c699](https://github.com/ignaciotitos/AutoIV/commit/3a5c69914382e0ff4853c232d184ce7706a0cad7)

## Objeto valor o entidad
Para cada clase ha sido necesario saber cómo va a funcionar cada clase, para saber si van a ser inmutables o mutables. Para esto ha sido necesario entender el problema para estructurar las clases. He propuesto una solución en el que las entidades sean Peticiones y Manejo y el objeto valor sea Empleado.
