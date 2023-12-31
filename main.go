package main

import (
	"flag"
	"fmt"
	"log"
)

func validate(target, file, wordlist string, stat, dir bool) {
	if (stat && dir) || (!stat && !dir) {
		log.Fatal("You have to choose one operation")
	}
	if dir && wordlist == "" {
		log.Fatal("You have to provide a wordlist for enumeration")
	}
	if (target == "" && file == "") || (target != "" && file != "") {
		log.Fatal("You have to provide either a single target or a list of targets")
	}
}

func exec(target, file, wordlist string, stat, dir bool, ch chan string, tlist []string) {
	validate(target, file, wordlist, stat, dir)
	if stat {
		if file != "" {
			tlist = readfile(file)
			for i := 0; i < len(tlist); i++ {
				go chckstat(tlist[i], ch)
			}
			for s := 0; s < len(tlist); s++ {
				fmt.Print(<-ch + "\n")
			}
		} else {
			chckstat(target, ch)
			fmt.Print(<-ch)
		}
	} else if dir {
		tlist = readfile(wordlist)
		for i := 0; i < len(tlist); i++ {
			go direnum(target, tlist[i], ch)
		}
		for s := 0; s < len(tlist); s++ {
			fmt.Print(<-ch + "\n")
		}
	}
}

func main() {
	var (
		file     string
		target   string
		wordlist string
		tlist    []string
		stat     bool
		dir      bool
	)
	ch := make(chan string, 100)
	flag.StringVar(&target, "t", "", "An ip or domain address to target")
	flag.StringVar(&file, "f", "", "File containing a list of targets (One target per line)")
	flag.StringVar(&wordlist, "w", "", "a Wordlist for enumeration (used with directory enumeration only)")
	flag.BoolVar(&stat, "s", false, "Check if there is a functioning website on target(s)")
	flag.BoolVar(&dir, "d", false, "Enumerate directories on target(s)")
	flag.Parse()

	exec(target, file, wordlist, stat, dir, ch, tlist)
}
