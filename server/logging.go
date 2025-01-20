package server

import (
	"log"
	"os"
)

func configureLogging() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("[INFO] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
