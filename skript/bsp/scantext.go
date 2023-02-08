//
//  scantext.go  --  Datei interpretieren
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"regexp"
)

var (
	r_colon = regexp.MustCompile(`:`)
	r_indent = regexp.MustCompile(`^\s+`)
)

func read_file() {
	file, err := os.Open("teams-go-3.vcs")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if r_indent.MatchString(line) {
			continue  // schlampig: Rest vergessen.
		}

		a := r_colon.Split(line, 2)
		field	 := strings.ToLower(a[0])
		contents := a[1]
		fmt.Printf("%-20s  --  %s\n", field, contents)


	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	read_file()
}

// vim: set noet ts=4 sw=4 sts=4 :
