// W GO _tablica_ jest sekwencją określonej i stałej ilości
// elementów. Zazwyczaj używamy [wycinków](slices), ale
// tablice bywają lepsze w niektórych przypadkach.

package main

import "fmt"

func main() {

	// Utworzymy tablicę `a` która może przechowywać dokładnie
	// pięć elementów typu `int`. Zarówno typ przechowywanych
	// elementów, jak i wielkość tablicy jest częścią typu.
	// Domyślnie tablica jest wypełniona elementami zerowymi
	// przypisanymi do typu, w tym przypadku `0` dla `int`.
	var a [5]int
	fmt.Println("emp:", a)

	// Możemy ustawić wartość w określonym miejscu tablicy,
	// używając składni `array[index] = value`. Jeśli
	// chcemy pobrać wartość, używamy `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Wbudowana funkcja `len` zwraca wielkość tablicy.
	fmt.Println("len:", len(a))

	// W ten sposób możemy zadeklarować i wypełnić tablicę
	// za jednym zamachem.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Tablice są jednowymiarowe, ale możemy użyć tablicy
	// jako przechowywanych wartości, tworząc strukturę
	// wielowymiarową.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
