package util

import "fmt"

// PrintHelp prints user help guide
func PrintHelp() {
	fmt.Println()
	for k, v := range HelpList {
		fmt.Printf("%s : %s\n", k, v)
	}
}
