package main

import (
	"RiskService/server"
	"log"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatalln("failed to start service due to error", err)
	}
}
