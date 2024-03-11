// Go umożliwia definiowanie _stałych_ typu: char, string, boolean,
// a także wartości numerycznych.

package main

import (
	"fmt"
	"math"
)

// `const` deklaruje wartość stałą.
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const` może być użyte wszędzie tam, gdzie `var`.
	const n = 500000000

	// Wyrażenia są wykonywane z określoną precyzją..
	const d = 3e20 / n
	fmt.Println(d)

	// Stałe liczbowe nie mają określonego typu, dopóki nie jest ona nadana
	// poprzez odpowiednią konwersję.
	fmt.Println(int64(d))

	// Stała liczbowa typ może nadać również kontekst,
	// w którym została ona użyta, jak przypisanie
	// bądź wywołanie funkcji. Dla przykładu: do
	// `math.Sin` powinniśmy przekazać `float64` (i ten typ
	// otrzymuje nasza zmienna w wywołaniu).
	fmt.Println(math.Sin(n))
}
