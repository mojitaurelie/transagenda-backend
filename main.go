package main

import (
	"fmt"
	"github.com/transagenda-back/server"
	"runtime"
)

type Version string

const version Version = "1.0.0"

func main() {
	fmt.Printf("Trans Agenda - Backend Server %s (%s %s)\n", version, runtime.GOOS, runtime.GOARCH)
	server.Serve()
}
