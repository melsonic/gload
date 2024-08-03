package main

import (
	"fmt"
	"os"

	"github.com/melsonic/gload/util"
)

// main
func main() {
	argData := util.GLoadArgs{RequestCount: 1, ConcurrentUsers: 1}
	argData.ParseArgs(os.Args[1:])
	if argData.HelpFlag {
		util.PrintHelp()
		return
	}

	// no error so do the main thing
	util.CoreFunction(&argData)

	fmt.Println("\n========== done ==========")
}
