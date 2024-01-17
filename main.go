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

func sortMapKeys(m []rune) []rune {
	sort.Slice(m, func(i, j int) bool {
		return m[i] < m[j]
	})
	return m
}

type Key struct {
	// TODO: match dinamically if using 3 rune
	// if only z exist, then z called instead of typing 3 rune 'zzz'
	// if only g g exist, then it called instead typing g g g

	// First, Second, Third rune
	First, Second rune
}

func main() {
	links := map[Key]string{
		{'g', 'n'}: "https://github.com/new",
		{'o', 'o'}: "https://onedrive.live.com/?id=FA1B916ECA036D35%213186&cid=FA1B916ECA036D35",
		{'2', '2'}: "https://mail.google.com/mail/u/5/",
		{'2', '1'}: "https://mail.google.com/mail/u/5/",
		{'1', '1'}: "https://mail.google.com/mail/u/0/#inbox",
		// '3':       "https://mail.google.com/mail/u/2/",
		// '4':       "https://mail.google.com/mail/u/3/",
		// 'l':       "https://lrc-maker.github.io/#/",
		// 't':       "https://tailwindcss.com/docs/installation",
		// 'g':       "https://gobyexample.com/",
		// 'n':       "https://www.lazyvim.org/keymaps",
		// 'a':       "https://aur.archlinux.org/",
		// 'm':       "https://www.myanonamouse.net/tor/browse.php?&tor%5BsrchIn%5D%5Btitle%5D=true&tor%5BsrchIn%5D%5Bauthor%5D=true&tor%5BsearchType%5D=all&tor%5BsearchIn%5D=torrents&tor%5Bcat%5D%5B%5D=53&tor%5Bcat%5D%5B%5D=75&tor%5Bcat%5D%5B%5D=0&tor%5BbrowseFlagsHideVsShow%5D=0&&&tor%5BminSize%5D=0&tor%5BmaxSize%5D=0&tor%5Bunit%5D=1&tor%5BminSeeders%5D=0&tor%5BmaxSeeders%5D=0&tor%5BminLeechers%5D=0&tor%5BmaxLeechers%5D=0&tor%5BminSnatched%5D=0&tor%5BmaxSnatched%5D=0&&tor%5BsortType%5D=default&tor%5BstartNumber%5D=0&thumbnail=true",
	}

	app := &cli.App{
		Name:  "goto-docs",
		Usage: "Open a link in your browser",
		Action: func(*cli.Context) error {
			var firstKeys []rune
			for key := range links {
				firstKeys = append(firstKeys, key.First)
			}
			// TODO: sort by Key struct by sum of First, Second, Third
			sortedKeys := sortMapKeys(firstKeys)
			for _, sortedKey := range sortedKeys {
				for key, link := range links {
					if sortedKey == key.First {
						fmt.Printf("%v %v - %v\n", string(key.First), string(key.Second), string(link))
					}
				}
			}

			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			char2, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}

			for key, link := range links {
				if char == key.First && char2 == key.Second {
					// fmt.Printf("%v %v - %v\n", string(key.First), string(key.Second), string(link))
					exec.Command("firefox", link).Run()
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
