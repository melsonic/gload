package util

// Command Line Argument Flags
var (
	UrlFlag             string = "-u"
	RequestCounterFlag  string = "-n"
	ConcurrentUsersFlag string = "-c"
	HelpFlag            string = "-h"
)

// Status Code
var SuccessStatusCode int = 200

// Command Line
var ParsingCLIArgsError string = "Passed Wrong Command Line Arguments"

// Help items
var HelpList map[string]string = map[string]string{
	UrlFlag:             "target url",
	RequestCounterFlag:  "total request count",
	ConcurrentUsersFlag: "concurrent user count",
}
