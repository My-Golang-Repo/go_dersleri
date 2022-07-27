package main

import (
	"os"
	"log"
)


var (
	newFile *os.File
	err error

)

func main() {
	newFile,err = os.Create("salam.txt")
	if err != nil {
		log.Fatal(err)
	}

}