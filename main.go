package main

import (
	"RiskService/server"
	"log"
)

func main() {
	s := server.NewServer()
	if s == nil {
		log.Fatalln("failed to create server")
		return
	}
	err := s.Run()
	if err != nil {
		log.Fatalln("failed to start server due to error", err)
		return
	}
}
