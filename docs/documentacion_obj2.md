# Documentación sobre la implementación
## Análisis del dominio
Después de tantos issues mal planteados puedo entender mejor a lo que el problema se refiere. Para esto un problema que me encontré es como se van a gestionar las peticiones [como se gestionarían las peticiones entre los empleados y las vacaciones](https://github.com/RafaelT00/AutoIV/issues/15), aunque el issue mal planteado también porque se tiene que discutir en el PR.

Más tarde apareció el problema del calendario según el M0, [como se va a gestionar entre las vacaciones y los empleados](https://github.com/RafaelT00/AutoIV/issues/16), que también se plante mal aunque queda resuelto el problema.

Después tuve que ver si había alguna [diferencia entre Empleado y el encargado de asignar vacaciones](https://github.com/RafaelT00/AutoIV/issues/17), en el cuál quedo claro que no había falta diferenciar entre estos dos.

## Objeto valor o entidad
Para cada clase ha sido necesario saber cómo va a funcionar cada clase, para saber si van a ser inmutables o mutables. Para esto ha sido necesario hablar con Rafael y entender cómo se van a asignar las distintas clases. Por ello, la clase Vacaciones, va a ser una entidad ya que van a ser mutables. Empleado en este caso se han hecho unas modificaciones para que las partes mutables de Empleado se tengan en cuenta al instanciar un empleado en Vacaciones, y desde esta clase poder ir modificando lo necesario para este. Empleado será el objeto valor.
