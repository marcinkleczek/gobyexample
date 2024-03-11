// Go posiada wiele typów pozwalających przechowywać napisy,
// liczby całkowite, rzeczywiste, wartości boolowskie itp.
// Oto kilka z nich.

package main

import "fmt"

func main() {

	// Stringi mogą być łączone znakiem `+`.
	fmt.Println("go" + "lang")

	// Liczby całkowite i zmiennoprzecinkowe.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Bool, z operatorami działającymi jak w innych językach.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
