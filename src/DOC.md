# Documentación sobre la implementación
## Análisis del dominio
Me surgió la necesidad de preguntar los atributos que iban a tener cada clase, Empleado, Vacaciones y Restricciones. 
- [Empleado](https://github.com/RafaelT00/AutoIV/issues/10)
- [Vacaciones](https://github.com/RafaelT00/AutoIV/issues/11)

Después de entender un poco más sobre las clases que iba a tener, tuve que preguntar como se iban a relacionar las clases entre ellas. Me surgió la duda de [como se gestionarían las peticiones entre los empleados y las vacaciones](https://github.com/RafaelT00/AutoIV/issues/15), en las que quedó claro que las tendrían que gestionarse desde los propios empleados.

Más tarde apareció la duda del calendario según el M0, [como se va a gestionar entre las vacaciones y los empleados](https://github.com/RafaelT00/AutoIV/issues/16), en el cuál según Rafael quedó bien definido gracias al issue de los atributos de Vacaciones.

Después tuve que ver si había alguna [diferencia entre Empleado y el encargado de asignar vacaciones](https://github.com/RafaelT00/AutoIV/issues/17), en el cuál quedo claro que no había falta diferenciar entre estos dos.

## Objeto valor o entidad
Para cada clase ha sido necesario saber cómo va a funcionar cada clase, para saber si van a ser inmutables o mutables. Para esto ha sido necesario hablar con Rafael y entender cómo se van a asignar las distintas clases. Por ello, las clases Restricciones y Vacaciones, van a ser entidades ya que van a ser mutables, mientras que Empleado que iba a ser inmutable de primeras, con los nuevos cambios que se hicieron añadiéndole variables que si van a cambiar durante el ciclo de vida de esta clase cuando se instancie, va a ser otra entidad. En este caso no hay objetos valor, no hay clases inmutables en el tiempo.

Así con Empleado se soluciona el hecho de crear empleados, ya que es necesario estos para asignar y que les asignen vacaciones. Con la clase Vacaciones se tiene un calendario de los días en los que se puede asignar vacaciones a los empleados, y con la clase Restricciones se soluciona el hecho de que no les puedes asignar vacaciones a los empleados a la vez, tienes que llevar un conteo y poder saber cuando asignar a un empleado o a otro.
