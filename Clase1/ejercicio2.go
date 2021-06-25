package main //Define el paquete main

// Importa librerías
import (
	"fmt"
)

func main() {

	var temperatura float32 //Medido en Centígrados
	var humedad int16       //Medidos en porcentaje
	var presion float32     //Medidos en atmósferas

	temperatura = 17.5
	humedad = 52
	presion = 1.1

	fmt.Println("Para la ciudad de Bogotá, se tiene los siguentes datos: ")
	fmt.Println("La temperatura es de ", temperatura, " grados centigrados")
	fmt.Println("La humedad es de ", humedad, " porciento")
	fmt.Println("La presion atmosférica es de ", presion, " atmósferas")

}
