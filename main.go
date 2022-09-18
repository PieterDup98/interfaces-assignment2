package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type sysFile struct {
	content string
}

func main() {
	var fileName string

	//Add CLI arg
	flag.StringVar(&fileName, "file", "go.txt", "The name of the file you want to read")
	flag.Parse()
	fmt.Println(os.Args[1:])

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error couldn't read file.", err)
	}

	var sf sysFile
	io.Copy(&sf, file)

	fmt.Println(sf.content)
}

func (sf *sysFile) Write(b []byte) (int, error) {
	fmt.Println("sysFile Write invoked []b size:", len(b))

	(*sf).content = string(b)
	return len(b), nil
}
