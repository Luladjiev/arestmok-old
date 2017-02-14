package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/luladjiev/arestmok/server"
)

func displayHelp() {
	fmt.Println("Arestmok is a tool for serving mocked data.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("\tarestmok command [arguments]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	fmt.Println("\tstart\tstart arestmok server")
	fmt.Println("\thelp\tdisplay this text")
	fmt.Println()
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	startCommand := flag.NewFlagSet("start", flag.ExitOnError)
	port := startCommand.Int("port", 8080, "A port on which the server will listen")

	if len(os.Args) < 2 {
		displayHelp()
	}

	switch os.Args[1] {
	case "start":
		startCommand.Parse(os.Args[2:])
	case "help":
		displayHelp()
	default:
		displayHelp()

	}

	switch true {
	case startCommand.Parsed():
		server.Run(*port)
	default:
		displayHelp()
	}
}
