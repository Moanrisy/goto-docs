package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"

	"github.com/urfave/cli"
)

func sortMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func main() {
	links := map[string]string{
		"t": "https://tailwindcss.com/docs/installation",
		"g": "https://gobyexample.com/",
		"l": "https://www.lazyvim.org/keymaps",
		"a": "https://aur.archlinux.org/",
	}

	app := &cli.App{
		Name:  "goto-docs",
		Usage: "Open a link in your browser",
		Action: func(*cli.Context) error {
			sortedKeys := sortMapKeys(links)
			for _, k := range sortedKeys {
				fmt.Println(k, links[k])
			}

			fmt.Print("Type a number: ")
			var i string
			fmt.Scanln(&i)

			exec.Command("firefox", links[i]).Run()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
