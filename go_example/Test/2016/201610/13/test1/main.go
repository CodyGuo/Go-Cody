package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	outFile, err := os.Create("test.zip")
	chekErr(err)
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"test.txt", "String contents of file."},
		{"test2.txt", "\x61\x62\x63\n"},
	}

	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		chekErr(err)

		_, err = fileWriter.Write([]byte(file.Body))
		chekErr(err)
	}
	err = zipWriter.Close()
	chekErr(err)
}

func chekErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
