package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"

	"github.com/eiannone/keyboard"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v3"
)

func sortMapKeys(m []rune) []rune {
	sort.Slice(m, func(i, j int) bool {
		return m[i] < m[j]
	})
	return m
}

type Key struct {
	// TODO: match dinamically if using 3 rune
	// if only z exist, then z instansly called instead of typing 3 rune 'zzz'
	// if only g g exist, then g g called instansly instead typing g g g

	// First, Second, Third rune
	First, Second rune
}

type Link struct {
	Name string
	Url  string
}

var data = `
config:
  gg:
    Name: Joko
    Url: https://google.com 
  go:
    Name: google ooo
    Url: https://google.comOOOO
`

func main() {
	m := make(map[interface{}]interface{})

	yamlFile, err := os.ReadFile("/mnt/storage2/Projects/practices/go/cli/goto-docs/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// err = yaml.Unmarshal([]byte(data), &m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m:\n%v\n\n", m)

	// Check if "config" key exists and is of the right type
	configValue, ok := m["config"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: 'config' key not found or not a map.")
		return
	}

	// Loop through the nested map under "config"
	links := map[Key]Link{}
	for key, value := range configValue {
		// fmt.Println("First:", string(key[0]))
		// fmt.Println("Second:", string(key[1]))
		// fmt.Println("Name:", value.(map[string]interface{})["Name"])
		// fmt.Println("Url:", value.(map[string]interface{})["Url"])
		// fmt.Println()

		k := Key{First: rune(key[0]), Second: rune(key[1])}
		n := value.(map[string]interface{})["Name"].(string)
		u := value.(map[string]interface{})["Url"].(string)
		l := Link{Name: n, Url: u}
		links[k] = l
	}

	// fmt.Println(links)

	// links := map[Key]Link{
	// 	{'g', 'n'}: {"github new repo", "https://github.com/new"},
	// 	{'o', 'o'}: {"onedrive skripsi", "https://onedrive.live.com/?id=FA1B916ECA036D35%213186&cid=FA1B916ECA036D35"},
	// 	{'2', '2'}: {"mail asamsulfat", "https://mail.google.com/mail/u/5/"},
	// 	{'1', '1'}: {"mail 1", "https://mail.google.com/mail/u/0/#inbox"},
	// 	{'3', '3'}: {"mail personal", "https://mail.google.com/mail/u/2/"},
	// 	{'4', '4'}: {"mail newsletter", "https://mail.google.com/mail/u/3/"},
	// 	{'l', 'l'}: {"lrc maker", "https://lrc-maker.github.io/#/"},
	// 	{'d', 't'}: {"docs tailwindcss", "https://tailwindcss.com/docs/installation"},
	// 	{'d', 'g'}: {"docs gobyexample", "https://gobyexample.com/"},
	// 	{'d', 'n'}: {"docs neovim with lazyvim keymaps", "https://www.lazyvim.org/keymaps"},
	// 	{'a', 'a'}: {"AUR", "https://aur.archlinux.org/"},
	// 	{'m', 'm'}: {"mam", "https://www.myanonamouse.net/tor/browse.php?&tor%5BsrchIn%5D%5Btitle%5D=true&tor%5BsrchIn%5D%5Bauthor%5D=true&tor%5BsearchType%5D=all&tor%5BsearchIn%5D=torrents&tor%5Bcat%5D%5B%5D=53&tor%5Bcat%5D%5B%5D=75&tor%5Bcat%5D%5B%5D=0&tor%5BbrowseFlagsHideVsShow%5D=0&&&tor%5BminSize%5D=0&tor%5BmaxSize%5D=0&tor%5Bunit%5D=1&tor%5BminSeeders%5D=0&tor%5BmaxSeeders%5D=0&tor%5BminLeechers%5D=0&tor%5BmaxLeechers%5D=0&tor%5BminSnatched%5D=0&tor%5BmaxSnatched%5D=0&&tor%5BsortType%5D=default&tor%5BstartNumber%5D=0&thumbnail=true"},
	// }

	app := &cli.App{
		Name:  "goto-docs",
		Usage: "Open a link in your browser",
		Action: func(*cli.Context) error {
			var firstKeys []rune
			for key := range links {
				firstKeys = append(firstKeys, key.First)
			}
			// TODO: sort by Key struct using multiplication of rune First, Second, Third instead of using only rune First
			sortedKeys := sortMapKeys(firstKeys)
			isPrinted := map[Key]bool{}
			for _, sortedKey := range sortedKeys {
				for key, link := range links {
					if sortedKey == key.First {
						if isPrinted[Key{key.First, key.Second}] {
							continue
						} else {
							fmt.Printf("%v %v -> %v\n", string(key.First), string(key.Second), string(link.Name))
						}

						isPrinted[Key{
							key.First, key.Second,
						}] = true

						break
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
					exec.Command("firefox", link.Url).Run()
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
