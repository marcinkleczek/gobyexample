// _Slices_ (wycinki) to ważns struktura danych. Pozwala tworzyć
// sekwencję podobnie do `array`, ale udostępnia znacznie lepszy
// interface obsługi danych.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// W przeciwieństwie do tablic wycinki są określona tylko
	// typem przechowywanych wartości (bez ilości elementów).
	// Niezainicjalizowane wycinek ma wartość nil
	// i ilość elementów (długość) równą 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// Żeby stworzyć pusty wycinek, o niezerowej długości
	// używamy wbudowanej funkcji `make`. Poniżej tworzymy
	// wycinek o długości 3 (wartość zerowa dla typu `string`
	// to ""). Domyślnie pojemność wycinka jest równa jego
	// długości. Jeśli zawczasu wiemy, że będziemy dodawali
	// elementy do wycinka, możemy podać dodatkowy parametr
	// pojemność do funkcji `make`.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// Pobieranie i ustawianie wartości odbywa się tak samo,
	// jak w przypadku tablic.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// Funkcja `len` zwraca ilość elementów.
	fmt.Println("len:", len(s))

	// Wycinki różnią się od tablic między innymi tym,
	// że możemy dodawać do nich wartości (tablice mają
	// stałą wielkość). Używamy do tego wbudowanej funkcji
	// `append` która zwraca wycinek zawierający przekazany wycinek
	// rozszerzony o pozostałe przekazane argumenty.
	// Zwróćmy uwagę, że musimy użyć zwróconej przez `append`
	// wartości, ponieważ może one zwrócić inny wycinek, niż
	// przekazany (zazwyczaj dzieje się tak, jeśli pojemność przekazanego
	// wycinka była zbyt mała).
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Wycinki mogą być również kopiowane przy użyciu wbudowanej
	// funkcji `copy`. Tworzymy pusty wycinek i kopiujemy do
	// niego zawartość wycinka `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Wycinki obsługują operator "slice" (części) o strukturze
	// `wycinek[od:do]`. Na przykład, dla pobrania części
	// składającej się z elementów `s[2]`, `s[3]` i `s[4]` użyjemy:
	l := s[2:5]
	fmt.Println("sl1:", l)

	// To wycinek od początku do (ale wyłączając!) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// Ten zaś od `s[2]` włącznie do ostatniego elementu wycinka.
	l = s[2:]
	fmt.Println("sl3:", l)

	// Możemy zadeklarować i wypełnić wycinek za jednym razem.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Pakiet `slices` zawiera sporo funkcji
	// pomocniczych, przydatnych przy pracy z wycinkami.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Wycinki mogą być (jak tablice) łączone w wielowymiarowe
	// struktury. Długość wewnętrznego wycinka może się zmieniać
	// (w tablicach musi być stała).
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
