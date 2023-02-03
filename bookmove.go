package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func scancopy(sdir, ddir string) string {
	fulllist := ""
	sourcedir := sdir
	// sourcedir, _ := os.Getwd()
	// root := "\\yourdocshere\\"
	filetypes := []string{".epub", ".pdf", ".docx"}
	// filetypes := []string{".epub", ".pdf", ".docx", ".txt"}

	for _, ftps := range filetypes {

		destdir := ddir + ftps[1:] + "\\"
		// destdir := sourcedir + root + ftps[1:] + "\\"
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
				fulllist = fulllist + "\n" + path
				fmt.Printf("\n  dest :  %s\n", ddir+ftps[1:]+"\\"+a)
				if path != ddir+ftps[1:]+"\\"+a {
					fcopy(path, ddir+ftps[1:]+"\\"+a)
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	return fulllist
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
