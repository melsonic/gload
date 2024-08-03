### A Load-Testing Tool in Golang 
A Load Tester is a tool to test a website performance when a batch of concurrent users tries to access it.

### Features
- add target url using `-u` option.
- customize total number of requests through `-n` option.
- customize concurrent users using `-c` option.

### Using the application
- first run `make build` or `go build -o gload main.go`
- next `./gload -u [URL] -n [REQUEST_COUNT] -c [CONCURRENT_USERS]`
- to fetch help flags, use `./gload -h`

### TODO
- Display requests/second 
- Display Total request time, time for first byte, time for last byte
