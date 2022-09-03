package main

import (
	"github.com/mrizkimaulidan/storial/internal/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}
