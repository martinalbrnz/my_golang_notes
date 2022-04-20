# Notas

## Correr programa
```
$ go run .
$ go help 
```
## Módulos de Go
Los módulos son packages de código que pueden descargarse y utilizarse (como en Node.js). Para hacer algo que ya fue implementado por otra persona, solo es necesario importar un modulo con
```
$ go get <package-name>
```
### fmt
Incluye funciones para formatear texto, mostrarlo en consola
### go.sum
En el archivo go.sum estan incluidos los hashes de las dependencias directas e indirectas del módulo.
