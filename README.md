# AutoIV
Prácticas IV
Problema:
El reparto de vacaciones y días libres en una empresa se puede convertir en una tarea bastante enrevesada. Existen muchos factores para decidir cuándo y a quién dárselas:
-Se suele dar prioridad a las personas que llevan más tiempo en la empresa que a los recien llegados.
-Los empleados con hijos suelen tener preferencia a la hora de elegir sobretodo en determinadas épocas del año cuando los niños no tienen clases.
-Por supuesto hay épocas en las que hay una mayor carga de trabajo en la empresa por lo que habrá pocos que se puedan ir.
-Hay que tomar en cuenta los periodos de los proyectos, cuántas personas necesitan estar disponibles para poder realizarlo, y aún más importante es qué personas concretas deben estar, como los líderes de los grupos.
-Puede pasar que dos personas no se puedan ir a la vez porque si no ya no es posible completar cierta tarea.
-Diferentes fases de un proyecto pueden tener menos o más carga de trabajo y se necesite más o menos personal.
-La incorporación a un proyecto ya empezado puede ser bastante complicada.
-Según la determinados departamentos pueden tener una mayor o una menor carga de trabajo según la época.
-Las personas que lleven más tiempo sin haber tomado días libres deberían tener más prioridad frente a las que se los tomaron más recientemente.
-Algunos empleados pueden tener razones personales de peso que deberían tomarse en cuenta.

Las empresas enviarían los perfiles de los empleados con todos los datos que se deberían tomer en cuenta para la elección, además de información sobre los proyectos que llevarán, las personas implicadas en ellos y las fechas importantes a tomar en cuenta. Dados todos estos datos, se realizará un primer borrador, que irá cambiando a medida que los empleados vayan eligiendo y aceptando los días que se les ha ofrecido hasta que todos hayan quedado conformes, en la medida de lo posible.


[HU01] Eustaquio Habichuela
  Eustaquio Habichuela, hombre de 40 años de edad, estudios secundarios y un dominio relativamente bueno de los ordenadores. Trabaja en la oficina de una empresa de 
  desarrollo de software y entre sus funciones está la de aprobar las peticiones de vacaciones que los empleados solicitan. Esta es una tarea que no solía tener 
  problemas pero en los últimos años la empresa ha crecido, lo que se traduce en un mayor número de empleados a revisar. Este aumento de empleados se debe también al 
  aumento de trabajos que la empresa realiza, lo que hace que haya que tener en cuenta los periodos de mayor carga. Además, hay que tener en cuenta que para ciertos 
  proyectos hay empleados que son estrictamente necesarios, ya sea para dirigirlos o para enseñar a las nuevas incorporaciones, etc.

  Todo ello hace que el la tarea de asignar las vacaciones y días libres se haya convertido para Eustaquio en un trabajo de mucha carga, lo que le interfiere en la 
  realización de sus otras funciones en la empresa, las cuales son de mayor importancia.

  Eustaquio necesita alguna herramienta que le ayude a agilizar este proceso de asignación y que le auxilie también en las posteriores solicitudes de cambios que los 
  empleados le puedan pedir.

1.Almacenar los datos de los empleados.
  Necesitamos almacenar los datos de los empleados. Únicamente almacenaremos los datos que vayan a influir en la decisión de que prioridad darle al empleado en la 
  elección de las vacaciones.
  Utilizaremos alguna estructura que almacene:
  -Nombre y Apellidos
  -Identificación
  -Posición en la empresa
  -Años en la empresa
  -A qué proyectos está asignado y el papel que desempeña en ellos
  -Si tiene hijos o no, y también si son pequeños de más grandes
  -Otros datos personales que deberían tomarse en cuenta
  Los dos primeros datos nos sirven para identificar al empleado y todos los demás son lo que utilizaremos para la toma de decisión. Para ello habrá que transformar 
  todos los datos a algún formato numérico que podamos utilizar para hacer los cálculos.

2.Almacenar el calendario de la empresa
  Necesitamos almacenar el calendario de trabajo de la empresa. Los proyectos tendrán un rango de tiempo dentro del calendario. Sobre este calendario es donde iremos 
  asignando los periodos de vacaciones.
  
  Necesita compilar para pasar al siguiente punto.
  
3.Asignación de los periodos
  Tenemos los perfiles de los empleados y tenemos el calendario sobre el que queremos realizar las asignaciones. El sistema tomará las solicitudes, comprobará que son 
  válidas e irá creando distintos borradores a medida le vayan llegando las solicitudes hasta que el empleado acepte la asignación recibida.
