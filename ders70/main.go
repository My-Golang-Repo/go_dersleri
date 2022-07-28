package main

import (
	"io"
	"log"
	"os"
)


func main() {
	file, err := os.Open("./dos1.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file)
	newfile, _ := os.Create("dos1_copy.txt")
	defer newfile.Close()

	//kopyalama
	byteWrite,_ := io.Copy(newfile,file)
	log.Printf("kopyalandi %d", byteWrite)
	rrr := newfile.Sync()
	if rrr != nil {
		log.Fatal(err)
	}
//----------------

	dosya, err2 := os.OpenFile("./dos.txt",os.O_RDONLY,0666)
	if err2 != nil {
		if os.IsPermission(err2){
			log.Println("Hata yazma izni reddedildi")
		}
	}

	dosya.Close()


}