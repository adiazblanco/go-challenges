package main

import "fmt"

/*
 * Reto #3
 * ¿ES UN NÚMERO PRIMO?
 * Fecha publicación enunciado: 17/01/22
 * Fecha publicación resolución: 27/03/22
 * Dificultad: MEDIA
 *
 * Enunciado: Escribe un programa que se encargue de comprobar si un número es o no primo.
 * Hecho esto, imprime los números primos entre 1 y 100.
 *
 * Información adicional:
 * - Usa el canal de nuestro discord (https://mouredev.com/discord) "🔁reto-semanal" para preguntas, dudas o prestar ayuda a la acomunidad.
 * - Puedes hacer un Fork del repo y una Pull Request al repo original para que veamos tu solución aportada.
 * - Revisaré el ejercicio en directo desde Twitch el lunes siguiente al de su publicación.
 * - Subiré una posible solución al ejercicio el lunes siguiente al de su publicación.
 * - Link del repo de retos (https://github.com/mouredev/Weekly-Challenge-2022-Kotlin/tree/main/app/src/main/java/com/mouredev/weeklychallenge2022)
 *
 */

func primo(value int) bool {
	if value <= 1 {
		return false
	}

	count := 0

	for i := 1; i <= value; i++ {
		if mod := value % i; mod == 0 {
			count++
		}

		if count == 2 && i < value {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Primo?")

	for i := 1; i <= 100; i++ {
		if primo(i) {
			fmt.Println(i, " es primo")
		}
	}
}
