/*This challenge is for getting username and password through command line
and comparing them with saved username and password to see if they match using
if-else condition */

package main

import (
	"fmt"
	"os"
)

func main() {
	username := []string{"festus", "racheal"}
	password := []string{"festusln", "rachealn"}

	if len(os.Args) < 3 {
		fmt.Println("Please ensure you enter both username and password")
		return
	}

	for i := 0; i < len(username); i++ {
		if os.Args[1] == username[i] && os.Args[2] == password[i] {
			fmt.Println("Logged In Successfully, Welcome", username[i])
			return
		}
	}
	fmt.Println("Invalid username or password")

}
