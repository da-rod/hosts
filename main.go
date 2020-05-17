package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

var (
	goPath = os.Getenv("GOPATH")
	in     = flag.String("sources", goPath+"/src/github.com/da-rod/hosts/sources.json", "file containing the sources to retrieve the lists")
	out    = flag.String("output", "/etc/unbound/unbound.conf.d/blacklist.conf", "output file name")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Read sources file
	f, err := os.Open(*in)
	quitOnErr(err)
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	quitOnErr(err)

	// Parse whitelist sources
	var w whitelist
	err = json.Unmarshal(data, &w)
	quitOnErr(err)

	// Parse blacklist sources
	var b blacklist
	err = json.Unmarshal(data, &b)
	quitOnErr(err)

	// Retrieve feeds and build lists
	allow := buildList(w.Whitelist)
	block := buildList(b.Blacklist)

	// Remove whitelisted domains from blacklist
	var hosts []string
	for domain := range block {
		if _, exists := allow[domain]; !exists {
			hosts = append(hosts, domain)
		}
	}
	sort.Strings(hosts)

	// Generate hosts file
	writeHostsFile(hosts)
}

func quitOnErr(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func writeHostsFile(domains []string) {
	out, err := os.Create(*out)
	quitOnErr(err)
	defer out.Close()

	w := bufio.NewWriter(out)
	for _, domain := range domains {
		w.WriteString(fmt.Sprintf("local-zone: %q redirect\n", domain))
		w.WriteString(fmt.Sprintf("local-data: \"%s A 0.0.0.0\"\n", domain))
	}
	err = w.Flush()
	quitOnErr(err)
}
