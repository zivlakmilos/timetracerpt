package main

import (
	"fmt"
	"os"

	"github.com/zivlakmilos/timetracerpt/timetrace"
)

func printUssage() {
	fmt.Printf("Ussage:\n")
	fmt.Printf("\tstatus - returns polybar status for timetrace\n")
}

func main() {
	if len(os.Args) == 1 {
		printUssage()
		return
	}

	if os.Args[1] == "status" {
		status, err := timetrace.CheckStatus()
		if err != nil {
			fmt.Printf("off\n")
			return
		}

		fmt.Printf("%s\n", status)
		return
	}
}
