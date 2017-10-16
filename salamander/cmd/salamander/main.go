package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/nasa9084/salamander/salamander"
)

type options struct {
	Listen string `short:"l" long:"listen" env:"SALAMANDER_LISTEN" default:":8080" description:"listening address"`
}

func main() { os.Exit(exec()) }

func exec() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Printf("%s", err)
		return 1
	}
	s := salamander.NewServer(salamander.ListenAddr(opts.Listen))
	log.Printf("server listening: %s", opts.Listen)
	if err := s.Run(); err != nil {
		log.Printf("%s", err)
		return 1
	}
	return 0
}
