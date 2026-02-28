package main

import "fmt"

func main() {
	fmt.Println("Mapy w Go")

	language := make(map[string]string) //tworzymy mapę o nazwie language, która jest mapą typu string i string

	language["JS"] = "JavaScript" //dodajemy element do mapy, klucz to "JS", a wartość to "JavaScript"
	language["PY"] = "Python"
	language["RB"] = "Ruby"

	fmt.Println("Lista wszystkich jezykow", language)
	fmt.Println("Jaki jezyk jest oznaczony jako JS?", language["JS"]) //daje nam to wartość dla klucza "JS"

	delete(language, "RB") //usuwamy element z mapy, klucz to "RB"
	fmt.Println("Lista wszystkich jezykow po usunieciu RB", language)

	//petle sa interesujace w golangu

	for key, value := range language { //iterujemy po mapie language, key to klucz, a value to wartość

		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
