// Kod warunkowy przy wykorzystaniu `if` i `else` jest prosty
// i nie różni się od innych języków. Musisz jednak pamiętać,
// że warunek musi zwracać wartość typu bool.

package main

import "fmt"

func main() {

	// Najprostszy przykład.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Możesz użyć `if` bez `else`.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Operatory logiczne `&&` (i) / `||` (lub) są często wykorzystywane
	// w instrukcjach warunkowych.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// Warunek możesz poprzedzić instrukcjami. Wszystkie zmienne
	// w nich zadeklarowane będą dostępne zarówno w `if` jak i
	// wszystkich kolejnych gałęziach.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Nie musisz (nie powinieneś!) używać nawiasów wokół warunku,
// ale klamerki są zawsze wymagane.
