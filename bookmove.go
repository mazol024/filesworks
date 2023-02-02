package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Hello")
	dstdir := "c:\\allbooks"
	os.MkdirAll(dstdir, 0777)

	err := filepath.Walk("c:/ebooks/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		switch {

		case !info.IsDir() && strings.Contains(path, ".epub"):
			a := path[strings.LastIndex(path, "\\")+1:]
			fmt.Printf(" name: %s\n", a)
			// fmt.Printf(" name: %s\n", path)
			// fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
			fcopy(path, dstdir+"\\"+a)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
func fcopy(src, dst string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Copied %d bytes.", bytesCopied)
}
