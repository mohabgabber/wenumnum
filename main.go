package main

import (
	"flag"
	"fmt"
	"log"
)

func validate(target, file string) {
	if file != "" && target != "" {
		log.Fatal("You can only provide a target or a file containing targets not both")
	}
	if file == "" && target == "" {
		log.Fatal("You didn't provide any target or file")
	}
}

func main() {
	file := ""
	target := ""
	quiet := false
	var stat bool
	var tlist []string
	flag.StringVar(&target, "t", "", "An ip or domain address to target")
	flag.StringVar(&file, "f", "", "File containing a list of targets (One target per line)")
	flag.BoolVar(&stat, "s", false, "Check if there is a functioning website on target(s)")
	flag.BoolVar(&quiet, "q", false, "Print results directly without header (Use if you wanna pipe the results)")
	flag.Parse()
	validate(target, file)

	if file != "" {
		tlist = readfile(file)
	}
	if !quiet {
		fmt.Println("This script provides multiple functionalities for enumeratiing web applications")
		if file != "" {
			fmt.Println("Enumerating targets from file: ", file)
			fmt.Println("Targets count: ", len(tlist))
		} else {
			fmt.Println("Targeting: ", target)
		}
		fmt.Println("---------------------------------------")
	}
	if stat {
		if file != "" {
			chckstat(target, tlist)
		} else {
			chckstat(target, tlist)
		}
	}
}
