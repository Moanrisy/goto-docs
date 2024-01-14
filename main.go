package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

// TODO: support other OS
// https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8

func main() {
	links := []string{"http://archlinux.org", "http://google.com", "http://facebook.com"}
	var i int

	app := &cli.App{
		Name:  "goto-docs",
		Usage: "Open a link in your browser",
		Action: func(*cli.Context) error {
			fmt.Print("Type a number: ")
			fmt.Scanln(&i)
			exec.Command("xdg-open", links[i]).Start()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
