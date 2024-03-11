// `switch` pozwala rozpatrywać wiele rozgałęzień jednego
// warunku.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Najprostszy `switch`.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Przy użyciu przecinków możesz rozdzielić kilka
	// wyrażeń obsługiwanych przez jeden kod. Możesz użyć opcjonalnego
	// przypadku `default` jeśli jest taka potrzeba.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` bez wyrażenia może być użyte jako alternatywna metoda
	// podobna logicznie do if/else. Wyrażenia `case`
	// nie muszą być stałymi (jak w przykładzie)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Poniższy `switch` porównuje typy zmiennej zamiast wartości.
	// Użyj tego wzorca, żeby poprawnie zinterpretować wartości
	// przesłane jako `interface{}`.
	// W tym przykładzie zmienna `t` zawiera typ odpowiedni
	// dla przekazanej do funkcji zmiennej.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
