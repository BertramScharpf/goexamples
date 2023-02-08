/*
 *  readline.go  --  Zeilen aus Datei lesen
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func read_stdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}


func read_file() {
	file, err := os.Open("/etc/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//read_stdin()  // Beende mit Strg-D
	read_file()
}

// vim: set noet ts=4 sw=4 sts=4 :
