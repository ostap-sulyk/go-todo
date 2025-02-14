package main

import "go-todo/api/server"

func main() {
	s := server.NewHTTPServer(3000)
	s.Start()
}
