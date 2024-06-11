package main

import (
	"go-commerce/app/server"
)

func main() {
	srv := server.NewServer()

	srv.Start()
}
