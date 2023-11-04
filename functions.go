package main

import (
	"bufio"
	"fmt"
	"log"
	http "net/http"
	"os"
)

func chckstat(target, file string) {
	if target != "" {
		resp, err := http.Get(target)
		if err != nil {
			fmt.Printf("Target: %v Does not exist", target)
			os.Exit(0)
		}
		fmt.Printf("Target: %v \t Status Code: %v\n", target, resp.StatusCode)
	} else {
		tlist := readfile(file)
		for _, c := range tlist {
			resp, err := http.Get(c)
			if err != nil {
				fmt.Printf("Target: %v Does not exist", c)
				os.Exit(0)
			}
			fmt.Printf("Target: %v \t Status Code: %v\n", c, resp.StatusCode)
		}
	}
}

func direnum(target, wordlist string) {
	lines := readfile(wordlist)
	for _, c := range lines {
		resp, _ := http.Get(target + "/" + c)
		if resp.StatusCode != 404 && resp.StatusCode != 301 {
			fmt.Printf("/%v \t -> %v\n", c, resp.StatusCode)
		}
	}
}

func readfile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}
