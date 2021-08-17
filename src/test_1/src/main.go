package main

import (
	"fmt"
)

func main() {

	// DECLARACION DE CONSTANTES
	const const1 float64 = 3.14
	const const2 = 3.14
	fmt.Println("const1", const1)
	fmt.Println("const2", const2)

	// DECLARCION DE VARIABLES ENTERAS
	var1 := 1
	var var2 int = 1
	var var3 int
	fmt.Println(var1, var2, var3)

	// ZERO VALUE
	var zero1 int
	var zero2 float64
	var zero3 string
	var zero4 bool
	fmt.Println(zero1, zero2, zero3, zero4)

	// CALCULO DE UNA ARE DEL CUADRADO
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado:", areaCuadrado)

	// Operacion aritmetica
	x := 10
	y := 50

	//SUMA
	result := x + y
	fmt.Println("suma", result)

	//RESTA
	result = y - x
	fmt.Println("resta", result)

	// MULTIPLICACION
	result = x * y
	fmt.Println("multiplicacion", result)

	// DIVICION
	result = y / x
	fmt.Println("divicion", result)

	// Divicion, modulo (remanente)
	result = y % x
	fmt.Println("Modulo", result)

	// INCREMENTAL
	x++
	fmt.Println("incremental", x)

	// DECREMENTAL
	x--
	fmt.Println("Decremental", x)

	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	// FMT
	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}
