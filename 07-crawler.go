package main

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"path"
	"fmt"
	"io"
)

const (
	url = "http://www.jdlingyu.fun/"
)

func main() {
	resp, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}
	var arr []string
	resp.Find(".pin-coat").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band, _ := s.Find("a.imageLink.image").Attr("href")
		arr = append(arr, band)
	})
	isPathOrCreate("./crawler_img")
	fmt.Println(arr)
	for _, v := range arr {
		resp, err := fetch(v)
		if err != nil {
			log.Fatal(err)
		}
		title := resp.Find("h2.main-title").Text()
		isPathOrCreate("./crawler_img/" + title)
		url := "./crawler_img/" + title
		resp.Find(".main-body").Find("a").Each(func(i int, s *goquery.Selection) {
			var band string
			var exists bool
			img := s.Find("img")
			band, exists = img.Attr("data-original")
			if !exists {
				band, _ = img.Attr("src")
			}
			title, _ := img.Attr("title")
			ext := path.Ext(band)
			filePath := url + "/" + title + ext
			newFile, _ := os.Create(filePath)
			fmt.Println(1, band)
			if band != "" {
				image, err := http.Get(band)
				if err != nil {
					log.Fatal(1111, err)
				}
				defer image.Body.Close()
				_, err = io.Copy(newFile, image.Body)
				if err != nil {
					log.Fatal(2222, err)
				}
			}
		})
	}
}

func fetch(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 判断状态是否成功
	if resp.StatusCode != 200 {
		log.Fatalf("请求失败", resp.StatusCode, resp.Status)
	}
	body, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func isPathOrCreate(url string) error {
	_, err := os.Stat(url)
	if err != nil || os.IsNotExist(err) {
		return os.Mkdir(url, os.ModePerm)
	}
	return nil
}