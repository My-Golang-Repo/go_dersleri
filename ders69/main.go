package main

import (
	"log"
	"os"
)

var (
	fileinfo os.FileInfo
	err error
)

func main() {
	fileinfo, err = os.Stat("../rename_file/dosya")
	if err != nil {
		if os.IsNotExist(err){
			log.Fatal("Dosya yok")
		}
		log.Fatal(err)
	}

	log.Fatal(fileinfo)

	
}
