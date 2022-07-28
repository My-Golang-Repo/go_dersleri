package main

import (
	"log"
	"os"
)


var(
	file *os.File
)

func main() {
	file,_ = os.OpenFile("demo2.txt",os.O_CREATE| os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()
/* 	
	file.WriteString(`dsdsd\n
	sada
	asd
	as
	dosyad`) */

	byteslice := []byte("bu dosyaya yazdiq")
	yazilmis,_ :=file.Write(byteslice)
	log.Printf("yazilmis byte %d \n", yazilmis)

}