# Zauważmy, że choć wycinki to inne typy niż tablice,
# funkcja `fmt.Println` wyświetla je tak samo.
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# Przeczytaj [ten świetny artykuł](https://go.dev/blog/slices-intro)
# napisany przez zespół developerów Go jeśli chcesz
# poznać szczegóły implementacyjne wycinków w Go.

# Teraz, kiedy poznaliśmy tablice i wycinki, przejdziemy
# do innej ważnej wbudowanej struktury danych: mapy.
