•	¿Qué tipo de dato necesitas para tener decimales en Go?
•	Float64
•	¿Cómo se convierte un int a float64?
•	
•	¿Qué función de fmt te permite controlar cuántos decimales mostrar?
•	%.2f, dependiendo de cuantos decimales se quiera se va cambiando el numero

Preguntas obligatorias
11.	¿Cuántas líneas tiene tu función main al final del taller? Cuéntalas con cuidado.
12.	41 líneas 

13.	Si tuvieras que agregar 5 operaciones más (raíz cuadrada, logaritmo, seno, coseno, módulo), ¿qué tan grande se haría tu main? ¿Sería fácil de leer para alguien que vea el código por primera vez?
Se agregarían 5 opciones mas al switch y se alargaría unas 20 líneas mas o menos en código, y yo creo que seguiría siendo muy entendible el código.
14.	Notaste que el código para 'pedir un número al usuario' o 'imprimir el resultado' se repite varias veces. ¿No sería mejor escribirlo una sola vez y reutilizarlo en muchos lugares?
15.	Lo de imprimir el resultado si me parece que se repite varias veces.
16.	Tu historial es una variable string gigante. ¿Qué pasaría si quisieras: ordenarlo alfabéticamente, eliminar la operación número 2, o contar cuántas veces se usó la operación de suma?
17.	
18.	Después de este taller, ¿qué fue lo más difícil de Go para ti? ¿Y lo más interesante?
19.	Me parece un lenguaje muy interesante , ya que a comparación de otros es algo mucho más simplificado

11.1 Comparación con tu experiencia previa de programación
Si ya programaste antes en Python, JavaScript, Java o C, compara honestamente con Go:
•	¿Qué te resultó más estricto en Go que en otros lenguajes? (Por ejemplo: la conversión de tipos, los imports no usados, las llaves obligatorias)
•	La conversión de tipos de datos
•	¿Qué te resultó más simple? (Por ejemplo: la estructura de switch, la concisión de := para declarar variables)
•	La estructura del switch está muy fácil de entender
•	¿Notaste algo que no existía en otros lenguajes que conoces? (Por ejemplo: el operador & para los punteros en Scan)
•	Si por ejemplo al momento de usar un printf y queremos usar un dato ya declarado se usa % y dependiendo que tipo sea se usa %d , %f , %v , etc.

11.2 Comparación entre el inicio y el final de tu propio código
Mira el código del Paso 1 y compáralo con el código final del Paso 5. Pregúntate:
•	¿Cuántas líneas tenía al inicio? ¿Cuántas líneas tiene ahora?
•	Al inicio tenia menos de 20 líneas y termino con 55 líneas de código
•	¿Sigue siendo igual de fácil de leer, o ya empieza a verse complicado?
•	Se ve un poco más complicado, pero sigue siendo entendible
•	Si tuvieras que entregárselo a otro estudiante para que lo entendiera sin ayuda, ¿lo entendería?
•	Si
•	¿Qué partes del código te parecen repetitivas? ¿Cuáles te gustaría 'simplificar' si supieras cómo?
•	El case ya hace un buen trabajo no creo que se pueda simplificar más que eso
