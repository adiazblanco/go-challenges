package main

import "fmt"

/*
 * Reto #2
 * LA SUCESIÓN DE FIBONACCI
 * Fecha publicación enunciado: 10/01/22
 * Fecha publicación resolución: 27/03/22
 * Dificultad: DIFÍCIL
 *
 * Enunciado: Escribe un programa que imprima los 50 primeros números de la sucesión de Fibonacci empezando en 0.
 * La serie Fibonacci se compone por una sucesión de números en la que el siguiente siempre es la suma de los dos anteriores.
 * 0, 1, 1, 2, 3, 5, 8, 13...
 *
 * Información adicional:
 * - Usa el canal de nuestro discord (https://mouredev.com/discord) "🔁reto-semanal" para preguntas, dudas o prestar ayuda a la acomunidad.
 * - Puedes hacer un Fork del repo y una Pull Request al repo original para que veamos tu solución aportada.
 * - Revisaré el ejercicio en directo desde Twitch el lunes siguiente al de su publicación.
 * - Subiré una posible solución al ejercicio el lunes siguiente al de su publicación.
 * - Link del repo de retos (https://github.com/mouredev/Weekly-Challenge-2022-Kotlin/tree/main/app/src/main/java/com/mouredev/weeklychallenge2022)
 *
 */

func main() {
	fmt.Println("Fibonacci")

	var prevFibonacci int
	fibonacci := 0

	for i := 1; i <= 50; i++ {
		fmt.Println(fibonacci)

		if fibonacci == 0 {
			prevFibonacci = fibonacci
			fibonacci = 1

		} else if prevFibonacci >= 1 {
			fibonacci = fibonacci + prevFibonacci
			prevFibonacci = fibonacci - prevFibonacci

		} else if fibonacci == 1 {
			prevFibonacci = fibonacci

		}
	}
}
