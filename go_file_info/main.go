package main

import (
	"log"
	"os"
)
var(
	fileinfo os.FileInfo
	err error
)
func main() {
	fileinfo, err := os.Stat("../go_file/salam.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(fileinfo.Mode()) 
	
}