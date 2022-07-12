# Notas

## Correr programa
```
$ go run .
$ go help 
```

[**Docs sobre cmd go**](https://pkg.go.dev/cmd/go)
## Módulos de Go
Los módulos son packages de código que pueden descargarse y utilizarse (como en Node.js). Para hacer algo que ya fue implementado por otra persona, solo es necesario importar un modulo con
```
$ go get <package-name>
```
### fmt
Incluye funciones para formatear texto, mostrarlo en consola
### go.sum
En el archivo go.sum estan incluidos los hashes de las dependencias directas e indirectas del módulo.

----------
### go mod init
```
go mod init example/hello

```

Es similar a cuando se inicializa un proyecto con npm init.
Crea un archivo go.mod con información del modulo y y las librerías requeridas. (como un package.json)

# Estructura del código

## package main
Al principio de nuestro archivo (hello.go), aparece `package main` que indica que el archivo debe ser compilado para ser ejecutable [**como explican aca**](https://thenewstack.io/understanding-golang-packages/). En el caso de que querramos que sea un package reutilizable, esta línea no debería ir (porque no sería necesario compilarlo para usar el código).

## import packages
Es el sector del código donde especificamos que packages vamos a necesitar importar para correr el código, en este caso, llamo al package "fmt", que me provee funcionalidades de I/O (Input/Output) por ejemplo para imprimir una linea en consola.

[**fmt docs**](https://pkg.go.dev/fmt)

```go
import (
	"fmt"

	"rsc.io/quote"
)
```

## func main()
Es la funcion principal (jaja) desde la que se llaman el resto de las funciones, y se ejecuta al correr el programa (como en C++, Dart, etc...), algo asi como un *entry point*.

```go
func main() {
	fmt.Println(quote.Go())
}
```

Fuera de main, se deben definir el resto de las funciones que luego quiera ejecutar, y llamarlas desde dentro de main, o no se podrá acceder a las mismas.