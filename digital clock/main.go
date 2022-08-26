/*Go code for digital or led clock display*/

package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen" //go get -u github.com/inancgumus/screen : ths is used to get this package
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"▇▇▇",
		"▇ ▇",
		"▇ ▇",
		"▇ ▇",
		"▇▇▇",
	}
	one := placeholder{
		"▇▇ ",
		" ▇ ",
		" ▇ ",
		" ▇ ",
		"▇▇▇",
	}
	two := placeholder{
		"▇▇▇",
		"  ▇",
		"▇▇▇",
		"▇  ",
		"▇▇▇",
	}
	three := placeholder{
		"▇▇▇",
		"  ▇",
		"▇▇▇",
		"  ▇",
		"▇▇▇",
	}
	four := placeholder{
		"▇ ▇",
		"▇ ▇",
		"▇▇▇",
		"  ▇",
		"  ▇",
	}
	five := placeholder{
		"▇▇▇",
		"▇  ",
		"▇▇▇",
		"  ▇",
		"▇▇▇",
	}
	six := placeholder{
		"▇▇▇",
		"▇  ",
		"▇▇▇",
		"▇ ▇",
		"▇▇▇",
	}

	seven := placeholder{
		"▇▇▇",
		"  ▇",
		"  ▇",
		"  ▇",
		"  ▇",
	}
	eight := placeholder{
		"▇▇▇",
		"▇ ▇",
		"▇▇▇",
		"▇ ▇",
		"▇▇▇",
	}

	nine := placeholder{
		"▇▇▇",
		"▇ ▇",
		"▇▇▇",
		"  ▇",
		"  ▇",
	}

	colon := placeholder{
		"   ",
		" ◼︎ ",
		"   ",
		" ◼︎ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear() //this clear the screen before printing new data
	for {
		screen.MoveTopLeft() //This move the cursor to the top left to ensure values are printing to same point
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		// fmt.Println(hour, min, sec)

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10], colon, digits[min/10], digits[min%10], colon, digits[sec/10], digits[sec%10],
		}
		for line := range clock[2] { //range clock[0] means range of 5 lines since each of the arry index is 5 lines. we can use other index like clock[2]
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 { //this condition is used to check the condition to animate the colon
					next = "   "
				}
				fmt.Print(next, " ", " ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second) //This is used to sleep the display so that we have one display per second
	}
}
