## Solución de problema del día de programador de Meli

Para solucionar el problema yo utilicé Golang y su librería default de imágenes. El problema fue solucionado con los siguientes pasos:

1. Se separa cada una de los cuadraditos del imagen y los pone en un array allSquares.

2. Encuentro el primero cuadradito del topo la izquierda y lo pongo en un array de solución llamado 'formedImage'

3. Hago un algoritmo de brute force para llenar todos los cuadraditos comparando los colores de sus cantos con los que ya estan en el image formado

4. Cuando todos los cuadraditos estan correctamente en sus posiciones el programa se termina

## Como ejecutar

Para el desarrollo yo utilizé la version 1.10 de Golang. Para ejecutar y generar la salida (archivo solution.png) debes ejecutar el comando `go run main.go square.go`.

Cualquier duda mandar mail a guilherme.moreira@mercadolivre.com