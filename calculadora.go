package main

import "fmt"

func main() {
	var a float64
	var b float64
	operaciones := 0
	//var suma float64 = a + b
	//var resta float64 = a - b
	//var multipliacion float64 = a * b
	//var division float64 = a / b

	fmt.Printf(`
	========== CALCULADORA v1 ==========`)
	fmt.Println("Ingrese el 1er numero:")
	fmt.Scanln(&a)

	fmt.Println("Ingrese el 2do numero:")
	fmt.Scanln(&b)

	fmt.Println("Seleccione la operacion a realizar(+, -, *, /,^,!):")
	fmt.Scanln(&operaciones)

	switch operaciones {
	case 1:
		fmt.Printf("La suma de %v + %v es: %v", a, b, a+b)
	case 2:
		fmt.Printf("La resta de %v - %v es: %v", a, b, a-b)
	case 3:
		fmt.Printf("La multiplicacion de %v * %v es: %v", a, b, a*b)
	case 4:
		if b != 0 {
			fmt.Printf("La division de %v / %v es: %v", a, b, a/b)
		} else {
			fmt.Println("Error: Division por cero")
		}
	case 5:
		resultado := 1.0
		for i := 1; i <= int(b); i++ {
			resultado *= a
		}
		fmt.Printf("La potencia de %v ^ %v es: %v", a, b, resultado)

	case 6:
		resultado := 1.0
		for i := 1; i <= int(a); i++ {
			resultado *= float64(i)
		}
		fmt.Printf("El factorial de %v es: %v", a, resultado)

	default:
		fmt.Println("Operacion no valida")
	}
}
