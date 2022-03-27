package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
 * Reto #1
 * ¿ES UN ANAGRAMA?
 * Fecha publicación enunciado: 03/01/22
 * Fecha publicación resolución: 27/03/22
 * Dificultad: MEDIA
 *
 * Enunciado: Escribe una función que reciba dos palabras (String) y retorne verdadero o falso (Boolean) según sean o no anagramas.
 * Un Anagrama consiste en formar una palabra reordenando TODAS las letras de otra palabra inicial.
 * NO hace falta comprobar que ambas palabras existan.
 * Dos palabras exactamente iguales no son anagrama.
 *
 * Información adicional:
 * - Usa el canal de nuestro discord (https://mouredev.com/discord) "🔁reto-semanal" para preguntas, dudas o prestar ayuda a la acomunidad.
 * - Puedes hacer un Fork del repo y una Pull Request al repo original para que veamos tu solución aportada.
 * - Revisaré el ejercicio en directo desde Twitch el lunes siguiente al de su publicación.
 * - Subiré una posible solución al ejercicio el lunes siguiente al de su publicación.
 * - Link del repo de retos (https://github.com/mouredev/Weekly-Challenge-2022-Kotlin/tree/main/app/src/main/java/com/mouredev/weeklychallenge2022)
 *
 */

func anagram(a string, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)

	fmt.Println(a, "<->", b)

	a = strings.ReplaceAll(a, " ", "")
	b = strings.ReplaceAll(b, " ", "")

	if a == b {
		return false
	}

	arrayA := strings.Split(a, "")             // declaration version 1
	var arrayB []string = strings.Split(b, "") //declaration version 2
	fmt.Println(arrayA, "<->", arrayB)

	sort.Strings(arrayA)
	sort.Strings(arrayB)
	fmt.Println(arrayA, "<->", arrayB)

	a = strings.Join(arrayA, "")
	b = strings.Join(arrayB, "")
	fmt.Println(a, "<->", b)

	return a == b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("It's an Anagram?")
	textA, _ := reader.ReadString('\n')

	fmt.Println("and?")
	textB, _ := reader.ReadString('\n')

	fmt.Println("Result: ", anagram(
		strings.ReplaceAll(textA, "\n", ""),
		strings.ReplaceAll(textB, "\n", "")))
}
