package basics

import "fmt"

var number int = 10             // variable global
var text string = "Hello World" // se puede definir una variable con var tanto dentro como fuera de la funcion
var boolean bool = true
var x, y, z int = 1, 2, 3      // definir multiples variables en una sola linea
var a, b, c = 1, "Hello", true // definir multiples variables en una sola linea infiriendo el tipo

func Variables() {
	fmt.Println(number, text, boolean)
	fmt.Println(x, y, z)
	fmt.Println(a, b, c)
	shortNumber := 1     // variable local
	shortText := "Hello" // solo se puede utilizar el operador := para definir variables dentro de una funcion
	shortBoolean := false
	sx, sy, sz := 1, 2, 3 // definir multiples variables en una sola linea sirve igual con :=
	sa, sb, sc := 1, "Hello", true
	fmt.Println(shortNumber, shortText, shortBoolean)
	fmt.Println(sx, sy, sz)
	fmt.Println(sa, sb, sc)
}
