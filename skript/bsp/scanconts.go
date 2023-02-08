//
//  scanconts.go  --  Datei interpretieren mit "line continuations"
//

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	r_indent = regexp.MustCompile(`^\s+(.*)`)
)

func read_file() {
	file, err := os.Open("teams-go-3.vcs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var entries []string
	for scanner.Scan() {
		if r_indent.MatchString(scanner.Text()) {
			parts := r_indent.FindStringSubmatch(scanner.Text())
			entries[len(entries)-1] = entries[len(entries)-1] + parts[1]
		} else {
			entries = append(entries, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Printf("%q\n", e)
	}
}

func main() {
	read_file()
}

// vim: set noet ts=4 sw=4 sts=4 :
