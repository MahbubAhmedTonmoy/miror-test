package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitlist = []string{"apple", "mango", "b"}

	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "c", "d")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	//memory make () new ()

	highscore := make([]int, 4)
	highscore[0] = 100
	highscore[1] = 90
	highscore[2] = 300
	highscore[3] = 120

	highscore = append(highscore, 400, 500)
	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))
	var languages = []string{"c", "c++", "java", "c#", "js", "go"}
	fmt.Println(languages)
	var deleteIndex int = 2
	languages = append(languages[:deleteIndex], languages[deleteIndex+1:]...)
	fmt.Println(languages)

	//maps
	programmigLanguage := make(map[string]int)

	programmigLanguage["c"] = 1
	programmigLanguage["c#"] = 2
	programmigLanguage["cPP"] = 3
	programmigLanguage["java"] = 4
	programmigLanguage["js"] = 5
	programmigLanguage["rubby"] = 6

	fmt.Println(programmigLanguage)
	
	delete(programmigLanguage, "c")

	fmt.Println(programmigLanguage)

	//loop in map
	for key, value := range programmigLanguage {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
