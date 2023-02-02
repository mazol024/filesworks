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
	sourcedir, _ := os.Getwd()
	root := "\\yourdocshere\\"
	filetypes := []string{".epub", ".pdf", ".docx"}
	// filetypes := []string{".epub", ".pdf", ".docx", ".txt"}

	for _, ftps := range filetypes {

		destdir := sourcedir + root + ftps[1:] + "\\"
		os.MkdirAll(destdir, 0777)

		err := filepath.Walk(sourcedir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}
			switch {

			case !info.IsDir() && strings.Contains(path, ftps):
				a := path[strings.LastIndex(path, "\\")+1:]
				// fmt.Printf(" \n name:\n %s\n", a)
				fmt.Printf("\n source : %s\n", path)
				fmt.Printf("\n  dest :  %s\n", sourcedir+root+ftps[1:]+"\\"+a)
				if path != sourcedir+root+ftps[1:]+"\\"+a {
					fcopy(path, sourcedir+root+ftps[1:]+"\\"+a)
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
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
