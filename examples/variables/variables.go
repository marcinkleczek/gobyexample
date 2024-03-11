// W go _zmienne_ są deklarowane wprost, a kompilator
// używa tych typów m.in. do sprawdzania poprawności wywołania
// funkcji.

package main

import "fmt"

func main() {

	// `var` deklaruje jedną lub więcej zmiennych.
	var a = "initial"
	fmt.Println(a)

	// Możesz zadeklarować wiele zmiennych za jednym razem.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go potrafi wywnioskować wyp zmiennej, jeśli nie jest podany.
	var d = true
	fmt.Println(d)

	// Zmienne zadeklarowane bez przypisania wartości są ustawione na
	// wartość zerową (specyficzną dla każdego typu). Dla przykładu:
	// wartość zerowa dla typu `int` to `0`.
	var e int
	fmt.Println(e)

	// Skrót `:=` pozwala zadeklarować zmienną i
	// przypisać do niej wartość. Jak tutaj na przykład
	// `var f string = "apple"`.
	// Można tak zrobić tylko wewnątrz funkcji.
	f := "apple"
	fmt.Println(f)
}
