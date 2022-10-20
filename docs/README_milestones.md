
Milestones

0.Almacenar los datos de los empleados.
  Necesitamos almacenar los datos de los empleados. Únicamente almacenaremos los datos que vayan a influir en la decisión de que prioridad darle al empleado en la elección de las vacaciones.
  Utilizaremos alguna estructura que almacene:
  -Nombre y Apellidos
  -Identificación
  -Posición en la empresa
  -Años en la empresa
  -A qué proyectos está asignado y el papel que desempeña en ellos
  -Si tiene hijos o no, y también si son pequeños de más grandes
  -Otros datos personales que deberían tomarse en cuenta
  Los dos primeros datos nos sirven para identificar al empleado y todos los demás son lo que utilizaremos para la toma de decisión. Para ello habrá que transformar todos los datos a algún formato numérico que podamos utilizar para hacer los cálculos.
  Es necesario que en este punto el programa compile y pueda acceder fácilmente a los datos almacenados.

1.Almacenar el calendario de la empresa
  Necesitamos almacenar el calendario de trabajo de la empresa en el periodo en el que se quiere realizar las asignaciones. Cada día habrá información de si hay en marcha un proyecto o una actividad. Sobre este calendario asignaremos las vacaciones de los empleados.
  Debe compilar correctamente y poder acceder a los datos para consultarlos y para añadir información sobre qué empleado/s están de vacaciones en ese periodo/día.
  
2.Asignación de los periodos
  La asignación se realizará sobre el calendario que previamente hemos creado. Cada empleado podrá tener varias opciones ordenadas por preferencia. Para cada solicitud se comprobará que durante el periodo solicitado realmente se pueda dar las vacaciones, es decir, que el empleado no sea estrictamente necesario en un determinado proyecto, tarea o cualquier otro motivo. Se comprobará entonces si en el calendario "hay hueco". Si no se pudiera se realiza la misma comprobación con las diferentes opciones que el empleado haya realizado.

  La asignación se realizará empezando desde la solicitud con mayor prioridad hasta las de menos.
  Los periodos solicitados deberán ser de una forma concreta, por ejemplo, solo se podrá solicitar periodos de x semanas.

  Puede que ocurrir que un empleado no tenga aceptada ninguna de sus opciones. No es esperable que se tengan todas las asignaciones a la primera, simplemente se irá notificando a los empelados para que vuelvan a realizar una solicitud, indicando además ahora los periodos disponibles. La nueva solicitud será necesario que se envíe en un corto plazo para poder continuar con las asignaciones, si no perderá su prioridad. Este proceso continuará hasta que se haya conseguido satisfacer todos los solicitantes.

  En este punto el programa deberá ser capaz de tomar las solicitudes y con la información disponible de los empleados rellenar el calendario sobre el que se quiera hacer las peticiones.
  
3. Cambios
  Es normal que después de haber acabado con las asignaciones haya reclamaciones o peticiones de cambios, por lo que debemos asegurarnos que sea fácilmente hacer cambios o intercambios entre empleados.
