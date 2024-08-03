package util

import (
	"fmt"
	"strconv"
)

// Struct to store the configs passed through args
type GLoadArgs struct {
	// target url
	Url string
	// total request to be made
	RequestCount int
	// number of concurrent users
	ConcurrentUsers int
	// user asked for help (y/n)
	HelpFlag bool
}

// prints the command line argument details
func (gla *GLoadArgs) PrintGLoadArgs() {
	fmt.Println("gla->Url   : ", gla.Url)
	fmt.Println("gla->count : ", gla.RequestCount)
	fmt.Println("gla->users : ", gla.ConcurrentUsers)
}

// Parse command line argument slice and set it to GLoadArgs object
func (gla *GLoadArgs) ParseArgs(args []string) {
	var argsLen int = len(args)
	var success bool = true
	var parsingError error
	var tempInt64Variable int64
	for i := range args {
		if args[i] == UrlFlag {
			if i+1 > argsLen {
				success = false
				break
			}
			gla.Url = args[i+1]
		} else if args[i] == RequestCounterFlag {
			if i+1 > argsLen {
				success = false
				break
			}
			tempInt64Variable, parsingError = strconv.ParseInt(args[i+1], 10, 64)
			if parsingError != nil {
				success = false
				break
			}
			gla.RequestCount = int(tempInt64Variable)
		} else if args[i] == ConcurrentUsersFlag {
			if i+1 > argsLen {
				success = false
				break
			}
			tempInt64Variable, parsingError = strconv.ParseInt(args[i+1], 10, 64)
			if parsingError != nil {
				success = false
				break
			}
			gla.ConcurrentUsers = int(tempInt64Variable)
		} else if args[i] == HelpFlag {
			gla.HelpFlag = true
			break
		}
	}
	if success == false {
		panic(ParsingCLIArgsError)
	}
}
