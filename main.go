package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

/*
func lenAndUpper(a string) (wordLength int, wordUpper string) {
	defer fmt.Println("done.")
	wordLength = len(a)
	wordUpper = strings.ToUpper(a)

	return

}
func repeat(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for i := 1; i <= 6; i++ {
		println(i)
	}

	for _, number := range numbers {
		total += number
	}
	println(total)
	return 1
}
func canIDrink(age int) bool {
	switch koreanAge := age; koreanAge {
	case 18:
		return true
	case 10:
		return true
	}
	return false

}
*/

func main() {
	jin := map[string]string{"name": "jin", "age": "20"}
	for key, value := range jin {
		fmt.Println(key, value)
	}
	favFood := []string{"taco", "ramen"}
	nico := person{favFood: favFood, name: "nico", age: 18}
	fmt.Println(nico)

}
