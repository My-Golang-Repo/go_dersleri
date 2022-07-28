package main

import (
	"log"
	"os"
)



func main() {
	filepath := "../go_file/salam.txt"
	newPath :="./dosya"
	err := os.Rename(filepath,newPath)
	if err != nil {
		log.Fatal(err)
	}
	
}