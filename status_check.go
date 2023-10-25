package main

import (
	"fmt"
	http "net/http"
	"os"
)

func chckstat(target string, file []string) {
	if target != "" {
		resp, err := http.Get(target)
		if err != nil {
			fmt.Printf("Target: %v Does not exist", target)
			os.Exit(0)
		}
		fmt.Printf("Target: %v \t Status Code: %v", target, resp.StatusCode)
	} else {
		for _, c := range file {
			resp, err := http.Get(c)
			if err != nil {
				fmt.Printf("Target: %v Does not exist", c)
				os.Exit(0)
			}
			fmt.Printf("Target: %v \t Status Code: %v\n", c, resp.StatusCode)
		}
	}
}
