package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	drive "google.golang.org/api/drive/v2"
)

var (
	fileId     = flag.String("fileid")
	outputPath = flag.String("output")
)

func main() {
	fmt.Println("Download Start.")
	client := &http.Client{}

	service, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
		return
	}

	if err := os.MkDirAll(filepath.Dir(*outputPath), 0755); err != nil {
		log.Fatalf("Unable to make Directory: %v", err)
		return
	}

	fd, err := os.Create(*outputPath)
	if err != nil {
		log.Fatalf("Unable to create File: %v", err)
		return
	}
	defer fd.Close()

	res, err := service.Files.Get(*fileId).Download()
	if err != nil {
		log.Fatalf("Unable to download File: %v", err)
		return
	}
	defer res.Body.Close()
	bytes, err := io.Copy(fd, res.Body)
	if err != nil {
		log.Fatalf("Unable to io Copy: %v", err)
		return
	}
	fmt.Println("Download Finish.")
}
