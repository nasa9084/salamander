package log

import (
	"log"
	"os"
)

var (
	// Info logger
	Info = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// Error logger
	Error = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
)
