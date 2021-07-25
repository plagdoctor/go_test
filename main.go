package main

import (
	"PlagDoctor/mydict"
	"fmt"
)

//	/home/miumiu/go/src/
func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	dictionary.Update(baseword, "Second")

	word, _ := dictionary.Search(baseword)
	fmt.Println(word)

	/*
		dictionary := mydict.Dictionary{"first": "First word"}
		word := "hello"
		definition := "Greeting"
		err := dictionary.Add(word, definition)
		if err != nil {
			fmt.Println(err)
		}
		hello, _ := dictionary.Search(word)
		fmt.Println(hello)
		err2 := dictionary.Add(word, definition)
		if err2 != nil {
			fmt.Println(err2)
		}

		err3 := dictionary.Delete(word)
		if err3 != nil {
			fmt.Println(err3)
		}

		fmt.Println(err3)
		hello, _ = dictionary.Search(word)

		fmt.Println(hello)
	*/
}
