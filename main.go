package main

import (
	"fmt"
	"github.com/transagenda-back/constant"
	"github.com/transagenda-back/server"
	"runtime"
)

func main() {
	fmt.Printf("Trans Agenda - Backend Server %s (%s %s)\n", constant.Version, runtime.GOOS, runtime.GOARCH)
	server.Serve()
}
