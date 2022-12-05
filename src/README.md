# Documentación sobre la implementación
## Análisis del dominio
Según el [Milestone 0](https://github.com/RafaelT00/AutoIV/milestone/6), para superar este hito es necesario crear una estructura de datos de empleados y otra para los proyectos y las restricciones. Además se necesita una estructura de datos de las vacaciones asignadas. Según la [HU01](https://github.com/RafaelT00/AutoIV/issues/7) se necesita asimilar las peticiones cuando son muchas y asignar vacaciones según esto.

Para esto se han creado issues sobre que estructuras son necesarias, y después de esto, saber que atributos tienen cada clase y además como están conectadas cada una entre las demás.

## Objeto valor o entidad
Para cada clase ha sido necesario saber cómo va a funcionar cada clase, para saber si van a ser inmutables o mutables. Para esto ha sido necesario hablar con Rafael y entender cómo se van a asignar las distintas clases. Por ello, las clases Restricciones y Vacaciones, van a ser objetos valor ya que van a ser mutables, mientras que empleado va a ser inmutable, por lo que por eso se va a tener como la entidad.
