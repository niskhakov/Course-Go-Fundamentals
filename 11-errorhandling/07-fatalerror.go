package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	/*
		... the Fatal functions call os.Exit(1) after writing the log message
	*/
}
