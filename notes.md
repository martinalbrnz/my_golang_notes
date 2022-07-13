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

---
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

---

# Variables y tipos de datos

## Declarar y asignar
Las variables pueden declararse y asignarse al mismo tiempo o ser declaradas primero y luego inicializadas.
Deben estar antecedidas por la palabra clave `var` y pueden declararse múltiples variables al mismo tiempo. Si las declaro sin asignar, debo especificar el `type`, que al ser de tipado estático, no cambiará durante el tiempo de ejecución.

Ejemplo de declaración y luego asignación:
```go
var myNum int // declare one var

var (
	myBool bool
	myFloat float32
	myString string
) // declare multiple vars

func main () {
	myBool = true // single initialization
	myFloat, myNum = 3.22, 5 // multiple initialization
	myString = "Hello world" // must have double quotes
}
```

En el caso de declarar y asignar a la vz, podemos usar `:=`, conocido tambien como *"walrus operator"*, que es una forma de asignación corta.

Tiene algunas reglas:

1. No puede usarse fuera de funciones, ya que la declaración debería estar precedida de una palabra clave:
```go
// no keywords, not possible
illegal := 40

// 'var' makes this legal
var myInt int = 40

func wal() {
	legal := 40 // because it's function scoped
}
```
2. No se pueden usar mas de una vez en el mismo scope, ya que el operador crea una nueva variable, y la misma ya esta declarada, por lo que arroja error.
```go
legal := 40
legal := 42 // This should throw error
```
3. Pueden usarse para asignar múltiples variables a la vez.
```go
foo, bar := 34, bool
tic, tac := "yes", "no"
```
4. Puede usarse varias veces en declaraciones múltiples siempre y cuando al menos una de las variables no esté previamente definida. 
En este caso estamos reasignando valores a las variables que ya tenían y creando las variables que no existían
```go
foo, bar := myFunc()
foo, cats := myFunc() // cats is new
bar, thing := myFunc() // thing is new
```
5. Se puede usar := para declarar una varible en un nuevo scope, y luego reasignar esa variable con =.
```go
var foo int = 40

func main() {
	foo := 70 // this is a new scope
	foo = 40 // here we are just reasigning
}
```
6. Se puede declarar el mismo nombre con := dentro de sentencias if, for, switch ya que las mismas pertenecen al scope de esa sentencia.
```go
foo := 40
if foo := myFunc(); foo == 314 {
	// here foo is scoped to 314
	// whatever the function does
}
// here foo still is 40
```

