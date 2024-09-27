package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	URL   = "https://storage.googleapis.com/panels-api/data/20240916/media-1a-i-p~s"
	DELAY = 50

	QUALITY = []string{"dhd", "dsd", "s", "wfs"}
)

type ImagerReg struct {
	Version int                          `json:"version"`
	Data    map[string]map[string]string `json:"data"`
}

func getImgname(quality, link, ext string) string {
	url, _ := url.Parse(link)

	noTilde := strings.ReplaceAll(url.Path[1:], "~", "_")
	noSlash := strings.ReplaceAll(noTilde, "/", "_")
	cleanedName := strings.Split(noSlash, ".")[0]

	return fmt.Sprintf("%s/%s%s", quality, cleanedName, ext)
}

func downloadImg(imgUrl, quality string) error {
	res, err := http.Get(imgUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch image: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil
	}

	ext, err := mime.ExtensionsByType(res.Header.Get("Content-Type"))
	if err != nil || len(ext) == 0 {
		return nil
	}

	fileName := getImgname(quality, imgUrl, ext[0])

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create new file in fs: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		return fmt.Errorf("failed to write image content to file: %w", err)
	}

	log.Printf("successfully downloaded image: %s", fileName)

	return nil
}

func main() {
	res, err := http.Get(URL)
	if err != nil {
		log.Fatalf("failed to fetch image registry: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read body: %v", err)
	}

	var imageReg ImagerReg
	if err := json.Unmarshal(body, &imageReg); err != nil {
		log.Fatalf("failed to parse body: %v", err)
	}

	for _, quality := range QUALITY {
		_, err := os.Stat(quality)
		if err != nil && os.IsNotExist(err) {
			os.Mkdir(quality, os.ModePerm)
		}
	}

	for _, imgData := range imageReg.Data {
		for _, quality := range QUALITY {

			if hd, ok := imgData[quality]; ok {
				if err := downloadImg(hd, quality); err != nil {
					log.Printf("failed to download %s version: %v\n", quality, err)
				}
			}

			time.Sleep(time.Duration(DELAY) * time.Millisecond)
		}
	}
}
