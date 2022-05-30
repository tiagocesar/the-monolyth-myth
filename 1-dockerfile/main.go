package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("This program will freeze on the line below until you press Ctrl + C...")
	<-exit

	fmt.Println("\nYou pressed Ctrl + C, so the program will now exit.")
}
