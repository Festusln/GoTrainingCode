//creating english turkish dictionary

package main

import (
	"fmt"
	"os"
)

func main() {
	/* english := []string{"good", "great", "perfect"}
	turkish := []string{"iyi", "harika", "mukemmel"}
	var flag int

	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Enter English word to search")
		return
	}
	query := args[0] //word to search
	for i, word := range english {
		if word == query {
			fmt.Printf("English: %s -> Turkish: %s\n", word, turkish[i])
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("Word not found")
	}
	*/
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Enter English word to search")
		return
	}
	query := args[0]
	dic := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mukemmel",
	}
	dic["up"] = "emilokan"
	dic["down"] = "egbe kinni yi wa"

	// fmt.Printf("%#v", dic)

	if value, ok := dic[query]; ok {
		fmt.Printf("Key: %q : value: %#v\n", query, value)
	}

}
