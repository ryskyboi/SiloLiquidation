package check

import "fmt"

func Check(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		panic(e)
	}
}
