package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type FileInfo struct {
	Name     string
	Size     int64
	Modified string
}

var fileList []FileInfo

func main() {
	// Reading the files from a directory
	files, err := os.ReadDir("./inputFiles")
	CheckError(err)
	for _, file := range files {
		fmt.Println(file)
	}

	// Copying the files into another directory
	source_dir := "./inputFiles/"
	destination_dir := "./processedFiles/"
	err = os.MkdirAll(destination_dir, 0755)
	CheckError(err)

	files, err = os.ReadDir(source_dir)
	CheckError(err)

	for _, file := range files {
		srcPath := source_dir + file.Name()
		dstPath := destination_dir + file.Name()
		err := copyFile(srcPath, dstPath)
		fmt.Println(file.Name() + "File copied to " + dstPath + " at " + time.Now().String())
		CheckError(err)

		jsonData, err := json.MarshalIndent(fileList, "", "  ")
		CheckError(err)
		err = os.WriteFile("files.json", jsonData, 0644)
		CheckError(err)
	}
	fmt.Println("Files copied Successfully")

	// Reading a file from the processed file
	file, err := os.Open("./processedFiles/" + files[0].Name())
	CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	os.RemoveAll("./processedFiles")
	fmt.Println("All Processed Files deleted")
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func copyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	CheckError(err)
	defer sourcefile.Close()
	destfile, err := os.Create(dest)
	CheckError(err)
	defer destfile.Close()
	_, err = io.Copy(destfile, sourcefile)
	CheckError(err)
	err = os.Chmod(dest, 0755)
	CheckError(err)
	err = os.Chtimes(dest, time.Now(), time.Now())
	CheckError(err)
	info, err := os.Stat(source)
	if err != nil {
		panic(err)
	}

	fileData := FileInfo{
		Name:     info.Name(),
		Size:     info.Size(),
		Modified: info.ModTime().Format(time.RFC3339),
	}

	fileList = append(fileList, fileData)
	return
}
