package main

import (
	"fmt"
	"runtime"
)

type Version string

const version Version = "1.0.0"

func main() {
	fmt.Printf("Trans Agenda - Backend Server %s (%s %s)\n", version, runtime.GOOS, runtime.GOARCH) 
}
