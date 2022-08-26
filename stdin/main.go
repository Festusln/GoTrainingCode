//Accepting input from user in terminal
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		in.Scan()
		fmt.Println("Scanned test is", in.Text())
	}

}
