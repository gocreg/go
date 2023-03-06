package main

import "fmt"

func main() {
	//Простейший вывод
	// fmt.Println("Hello, world!")
	// fmt.Println("Привет, мир!")
	//форматированный вывод
	//fmt.Printf("\nПривет %s\nмне %d лет", "Игорь", 49)

	var name string = "Igor"
	var age int = 49

	age, name, height := 51, "Фоммичёв", 173
	fmt.Printf("Меня зовут %s, мне %d лет мой рост %d", name, age, height)
}
