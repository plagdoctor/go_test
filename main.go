package main

import (
	"PlagDoctor/mydict"
	"fmt"
)

//	/home/miumiu/go/src/
func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		println(definition)
	}

}
