# Documentación sobre la implementación
## Análisis del dominio
Me surgió la necesidad de preguntar los atributos que iban a tener cada clase, Empleado, Vacaciones y Restricciones. 
- [Empleado](https://github.com/RafaelT00/AutoIV/issues/10)
- [Vacaciones](https://github.com/RafaelT00/AutoIV/issues/11)

Después de entender un poco más sobre las clases que iba a tener, tuve que preguntar como se iban a relacionar las clases entre ellas. Me surgió la duda de [como se gestionarían las peticiones entre los empleados y las vacaciones](https://github.com/RafaelT00/AutoIV/issues/15), en las que quedó claro que las tendrían que gestionarse desde los propios empleados.

Más tarde apareció la duda del calendario según el M0, [como se va a gestionar entre las vacaciones y los empleados](https://github.com/RafaelT00/AutoIV/issues/16), en el cuál según Rafael quedó bien definido gracias al issue de los atributos de Vacaciones.

Después tuve que ver si había alguna [diferencia entre Empleado y el encargado de asignar vacaciones](https://github.com/RafaelT00/AutoIV/issues/17), en el cuál quedo claro que no había falta diferenciar entre estos dos.

## Objeto valor o entidad
Para cada clase ha sido necesario saber cómo va a funcionar cada clase, para saber si van a ser inmutables o mutables. Para esto ha sido necesario hablar con Rafael y entender cómo se van a asignar las distintas clases. Por ello, la clase Vacaciones, va a ser una entidad ya que van a ser mutables. Empleado en este caso se han hecho unas modificaciones para que las partes mutables de Empleado se tengan en cuenta al instanciar un empleado en Vacaciones, y desde esta clase poder ir modificando lo necesario para este. Empleado será el objeto valor.
