package main

import (
	"fmt"
	"time"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    sigs := make(chan os.Signal)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {  // Vorsicht: Funktioniert nicht sauber.
		for {
			sig := <-sigs
			fmt.Println()
			fmt.Println(sig)
		}
    }()

	time.Sleep(60 * time.Second)
}
