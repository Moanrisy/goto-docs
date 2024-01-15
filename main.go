package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"

	"github.com/eiannone/keyboard"
	"github.com/urfave/cli"
)

func sortMapKeys(m map[rune]string) []rune {
	keys := make([]rune, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func main() {
	links := map[rune]string{
		't': "https://tailwindcss.com/docs/installation",
		'g': "https://gobyexample.com/",
		'l': "https://www.lazyvim.org/keymaps",
		'a': "https://aur.archlinux.org/",
		'm': "https://www.myanonamouse.net/tor/browse.php?&tor%5BsrchIn%5D%5Btitle%5D=true&tor%5BsrchIn%5D%5Bauthor%5D=true&tor%5BsearchType%5D=all&tor%5BsearchIn%5D=torrents&tor%5Bcat%5D%5B%5D=53&tor%5Bcat%5D%5B%5D=75&tor%5Bcat%5D%5B%5D=0&tor%5BbrowseFlagsHideVsShow%5D=0&&&tor%5BminSize%5D=0&tor%5BmaxSize%5D=0&tor%5Bunit%5D=1&tor%5BminSeeders%5D=0&tor%5BmaxSeeders%5D=0&tor%5BminLeechers%5D=0&tor%5BmaxLeechers%5D=0&tor%5BminSnatched%5D=0&tor%5BmaxSnatched%5D=0&&tor%5BsortType%5D=default&tor%5BstartNumber%5D=0&thumbnail=true",
	}

	app := &cli.App{
		Name:  "goto-docs",
		Usage: "Open a link in your browser",
		Action: func(*cli.Context) error {
			sortedKeys := sortMapKeys(links)
			for _, k := range sortedKeys {
				fmt.Println(string(k), links[k])
			}

			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			fmt.Printf("You pressed: %q\r\n", char)

			exec.Command("firefox", links[char]).Run()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
