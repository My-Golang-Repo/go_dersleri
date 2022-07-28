package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
func main(){
	tempDir, _ := ioutil.TempDir("","gecici")
	fmt.Println("Temp dir: \n", tempDir)

	tempfile, _:= ioutil.TempFile(tempDir, "gecivi file")
	fmt.Printf("Temp file %v\n", tempfile.Name())

	_=os.Remove(tempfile.Name())
	_=os.Remove(tempDir)
	
	//fmt.Printf("Temp file %v\n", tempfile.Name())
	_ = os.Remove("ddd")
}