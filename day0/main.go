package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func looop(numbers ...int) int {
	for i := 1; i < len(numbers); i++ {
		fmt.Println(i)
	}
	return 1
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func swap(num1, num2 int) (Nnum1, Nnum2 int) {
	tmp := num1
	Nnum1 = num2
	Nnum2 = tmp
	return
}

func lenAndName(name string) (lenght int, UpperName string) {
	lenght = len(name)
	UpperName = strings.ToUpper(name)
	return
}

func mulitply(a, b int) int {
	return a * b
}

func main() {
	totalLen, Name := lenAndName("SongKitae")
	fmt.Println(totalLen, Name)

	num1, num2 := swap(1, 2)
	fmt.Println(num1, num2)

	looop(1, 2, 3, 4, 5)

	fmt.Println(mulitply(2, 2))

	fmt.Println(superAdd(1, 2, 3, 4, 5, 6))

	fmt.Println(canIDrink(19))

	arr := [5]string{"song", "nico", "lynn", "dal"}
	arr[4] = "123"
	fmt.Println(arr)

	slice := []string{"123", "123123", "45", "51235412", "1010"}
	slice = append(slice, "nico")
	fmt.Println(slice)

	maps := map[string]string{"name": "nico", "age": "12"}
	for key, value := range maps {
		fmt.Println(key, value)
	}

	favFood := []string{"kimchi", "ramen"}
	song := person{name: "Song Kitae", age: 17, favFood: favFood}
	fmt.Println(song.name)
}
