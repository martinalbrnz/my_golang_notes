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

# Operadores y palabras clave
### Identificadores
Los identificadores son secuencias de letras y/o digitos que identifican entidades de un programa, como variables o tipos.
por ejemplo:
```go
var holaMundo = "Hola mundo!"
```
En este caso holaMundo es un identificador
### Palabras clave (keywords)
Las palabras clave estan reservadas por el lenguaje y no pueden usarse como identificadores.
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Operadores
```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```
Muchos de estos son los típicos operadores de otros lenguajes.
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

## Tipos
Go incluye varios types, numéricos, strings, booleanos, errores, y la capacidad para crear tipos personalizados.
### Number
- Las variables numéricas tienen variantes según el espacio en memoria con 8, 16, 32 y 64 bit, pudiendo ser con signo (int) o sin signo (uint).
Un byte es un alias para un uint8 (entero sin signo de 8 bit). Una runa (rune) es un alias para un int32 (entero de 32 bit con signo). 
- Los números de coma flotante son float32 o float64.
- Soporta también números complejos, representados como complex32 o complex64.
### String
Los string son una secuencia de caracteres UTF-8, encerrados en comillas dobles.
Son inmutables, una vez creados no se puede cambiar el contenido del mismo. Podemos acceder al largo del string con la funcion len().
No puedo acceder a la locación en memoria de un elemento de un string, por ejemplo, si tengo `s[i]`, no puedo acceder a `&s[i]`.
Por defecto, toma el valor "".

### Arrays
Se pueden guardar listas de elementos en arrays, slices o maps
Los **array** se definen con un tamaño fijo y un tipo de dato comun para todos sus elementos.
Al ser el tamaño parte de la definición, los array no pueden crecer ni achicarse (para eso es necesario otro type)
```go
var  catNames = [4]string{"Pol", "Snow", "Daisy", "Orlando"}
```
Siempre una sola dimension, aunque pueden estar formados por otros arrays para formar estructuras multidimensionales.

### Slices
Los **slices** son similares a los slices de python. Actuan sobre un array y el numero de elementos visibles del slice determina su largo (length).

Si el Array sobre el que actua el slice es mayor que el slice, este último todavía tiene la capacidad de crecer.
Podríamos pensar en length como la cantidad actual de elementos y en capacity como el número máximo de elementos posibles.
ej: `var someCats = catNames[1:3]` incluye los catNames en los índices 1 y 2, no así el del índice 3.

Es importante destacar que si intentamos, por ejemplo, pasar `var someCats = catNames[1:5]` a catNames, cuyo length es 4, vamos a obtener el error *"invalid argument: index 5 (constant of type int) is out of bounds "*.

Pueden crearse slices de la misma forma que un array, pero sin especificar el tamaño, de forma de poder aplicarle métodos de slices.

```go
var intSlice = []int16{1, 2, 3, 4, 5, 6, 7}
```

Las posibilidades de los slice son similares a las de python, por ejemplo:

```go
var allCats = catNames[:] // copy of catNames array
allCats = append(allCats, "Sparkle") // append "Sparkle" in the end, not possible with arrays
var allCats = catNames[:2] // elements on index 0, 1
var allCats = catNames[2:] // elements on index 2, 3, 4...
```
Al igual que los array, solo pueden tener una dimensión, pero también pueden tener como elementos otros array.
Considerar que estos slices internos deben inicializarse por separado.

### Structs
Los structs son una secuencia de elementos nombrados, cada uno de los cuales tiene un nombre único y un tipo.
```go
// empty struct
struct {}

// struct with 6 fields
struct {
  x, y int
  u float32
  _ float32
  A *[]int
  F func()
} 
```
Un campo declarado con un tipo (*type*) pero sin nombre (*name*) se denomina campo incrustado (*embedded field*).
Este debe ser especificado con un nombre de tipo `T` o como un puntero (*pointer*) a un nombre de tipo no interface `*T` y T en si mismo no debe ser de tipo pointer
Un campo o método `f` de un campo incrustado en un struct x se denomina com promovido (*promoted*) si `x.f` es un selector legal que denota ese campo o método f.
Estos actuan como campos normales de un struct con la diferencia de que no pueden ser usados como *field names* (algo asi como privados? idk...).

### Pointers
Denota un set (lista de elementos únicos) de todos los pointers a variables de un tipo dado.
```go
*Point
*[4]Point
```
Los punteros hacen referencia a una dirección en memoria donde esta guardado determinado valor, no al valor en si mismo.
```go
// No cambia el valor de x
func zero(x int) {
  x = 0
}

// Recibe una dirección de memoria como parametro
func otherZero(xPtr *int) {
  // Cambia el valor de esa dirección
  *xPtr = 0
}

func main() {
  x := 5
  zero(x)
  fmt.Println(x) // Sigue siendo 5

  otherZero(&x) // Pasamos la dirección en memoria
  fmt.Println(x) // ahora x es 0
}
```
Se puede ver la diferencia entre los dos operadores * y &.
En este caso, primero representamos el puntero con `*int`. También podemos usar el * para desreferenciar las variables del puntero, por ejemplo: 
```go
*xPtr = 0 // Guardar int 0 en esta dirección
xPtr = 0 // Error! No es un int, es *int
``` 

El operador & nos sirve para referirnos a la dirección de una variable, por eso cuando llamamos a otherZero() le pasamos &x como parámetro.
En este caso &x retorna un valor de tipo *int, ya que x es un int.
&x en main() y xPtr en otherZero hacen referencia a la misma dirección de memoria.

También puede obtenerse un pointer con la función new(), que toma un tipo como argumento, asigna una cantidad suficiente de memoria para guardar un objeto de ese tipo y devuelve un pointer a esa dirección.
```go
xPtr := new(int)
```

### Maps
Es una estructura de datos que almacena pares key-value que tienen un tiempo de búsqueda constante, a costas de aleatorizar el orden de los pares key-value.

```go
catsWeight := map[string]float64{
  "Pol": 4.2, "Snow":  5.1, "Daisy": 3.8, "Orlando": 5.2,
}

for k, v := range catsWeight {
  fmt.Println(k, "weight is about", v, "k")
}
```

Notar como puedo acceder a key y a value por separado con k y v (similar a python)

