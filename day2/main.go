package main

import (
	"fmt"
	"learngo/day2/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First Word"}
	baseWord := "hello"
	definition := "Greeting"

	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// dictionary.Add(baseWord, definition)
	// err := dictionary.Update(baseWord, "Hi~")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	// hello, _ := dictionary.Search(baseWord)
	// fmt.Println("found", baseWord, "definition", hello)
	// err2 := dictionary.Add(baseWord, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	dictionary.Add(baseWord, definition)
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
