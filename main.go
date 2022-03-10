package main

import (
	"errors"
	"go_linux/log"
)

var lorn log.Linux_logger

func main() {
	//name := flag.String("n", "file.go", "Enter the name of the file as an argument")
	//author := flag.String("a", "Ayobami", "Enter the author of the file as an argument")
	//all := flag.String("all", "", "Get all of the files created so far")
	//flag.Parse()
	//
	////implement the native linux argument style using cobra
	//var null string
	//null = ""
	//
	//if *all != null {
	//	_, err := cmd.ListFilesCreated()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
	//
	//if *author != null && *name != null {
	//	var f cmd.File
	//	Content := "module go_linux\n\ngo 1.17"
	//	err := f.CreateFile(*name, *author, Content)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	testerr := errors.New("Test Error: Check blah blah")
	log.Log(testerr)
}
