package main

import "fmt"

/*
 * Reto #0
 * EL FAMOSO "FIZZ BUZZ"
 * Fecha publicación enunciado: 27/12/21
 * Fecha publicación resolución: 19/03/22
 * Dificultad: FÁCIL
 * Enunciado: Escribe un programa que muestre por consola (con un print) los números de 1 a 100 (ambos incluidos y con un salto de línea entre cada impresión), sustituyendo los siguientes:
 * - Múltiplos de 3 por la palabra "fizz".
 * - Múltiplos de 5 por la palabra "buzz".
 * - Múltiplos de 3 y de 5 a la vez por la palabra "fizzbuzz".
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
	fmt.Println("FIZZ BUZZ Challenge!")

	for i := 1; i <= 100; i++ {

		if mod3_5 := (i % 3) + (i % 5); mod3_5 == 0 {

			fmt.Println("FizzBuzz")

		} else if mod3 := i % 3; mod3 == 0 {

			fmt.Println("Fizz")

		} else if mod5 := i % 5; mod5 == 0 {

			fmt.Println("Buzz")

		} else {

			fmt.Println(i)
		}
	}
}
