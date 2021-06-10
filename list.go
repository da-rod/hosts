package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func buildList(s sources) map[string]bool {
	list := make(map[string]bool)

	for i := 0; i < len(s); i++ {
		for _, url := range s[i].URLs {
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			defer res.Body.Close()
			scanner := bufio.NewScanner(res.Body)
			for scanner.Scan() {
				l := scanner.Text()
				if strings.HasPrefix(l, "#") || len(l) == 0 || existsInStringSlice(l, s[i].Exclude) {
					continue
				}
				l = strings.ToLower(l)
				if s[i].Format == "hosts" {
					if strings.HasPrefix(l, "0.0.0.0") || strings.HasPrefix(l, "127.0.0.1") {
						l = strings.Split(l, " ")[1]
					} else {
						continue
					}
				}
				if _, exists := list[l]; exists {
					continue
				}
				list[l] = true
			}
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
		}
	}

	return list
}

func existsInStringSlice(s string, sl []string) bool {
	for _, entry := range sl {
		if s == entry {
			return true
		}
	}
	return false
}
