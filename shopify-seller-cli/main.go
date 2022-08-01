package main

import (
	"fmt"
	"os"
)

func main() {
	passedArgs := os.Args[1:]
	for i, cArg := range passedArgs {
		if len(passedArgs) > 3 {
			fmt.Println("Too many arguments provided. Exiting... ")
			os.Exit(0)
		} else if i < 2 {
			if cArg == "About" {
				fmt.Println("Some information about this program")
			} else if cArg == "Help" {

			}
		}
	}
	//fmt.Println(os.Args[1:])
}
