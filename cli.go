package main

import (
	"flag"
	"fmt"
	"os"
)

func cli() {

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	setCmd := flag.NewFlagSet("set", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'set' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "get":
		getCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'get'")

		fmt.Println("addr: ", os.Args[2])
		fmt.Println("data: ", os.Args[3])
	case "set":
		setCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'set'")
		fmt.Println("  tail:", setCmd.Args())
	default:
		fmt.Println("expected 'get' or 'set' subcommands")
		os.Exit(1)
	}
}
