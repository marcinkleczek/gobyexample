// `for` to jedyna konstrukcja pętli w Go. Poniżej są przykłady
// zastosowania tejże pętli.

package main

import "fmt"

func main() {

	// Pojedynczy warunek.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Klasyczna pętla `for`: ustawienie/warunek/wykonaj_po.
	for j := 0; j <= 3; j++ {
		fmt.Println(j)
	}

	// Inny sposób na osiągnięcie "zrób to N razy"
	// to iteracja `range` przez liczby.
	for i := range 3 {
		fmt.Println("range", i)
	}

	// `for` bez warunku będzie pętlą nieskończoną,
	// dopóki nie przerwiesz jej `break` lub `return`
	// które przerywa całą funkcję.
	for {
		fmt.Println("loop")
		break
	}

	// Możesz również pominąć część instrukcji w pętli,
	// używając `continue` przjdziesz do kolejnego wykonania pętli.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
