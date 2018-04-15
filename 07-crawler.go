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

type Album struct {
	title string
	images []Image
	resp *goquery.Document
}

type Image struct {
	name string
	url string
	ext string
}

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
		al := Album{title: title, resp: resp}
		al.getAllImage()
		for _, v := range al.images {
			v.download("./crawler_img/" + al.title)
		}
	}
}

func (al *Album) getAllImage() {
	al.resp.Find(".main-body").Find("a").Each(func(i int, s *goquery.Selection) {
		var band string
		var exists bool
		imgDom := s.Find("img")
		band, exists = imgDom.Attr("data-original")
		if !exists {
			band, _ = imgDom.Attr("src")
		}
		title, _ := imgDom.Attr("title")
		ext := path.Ext(band)
		if band != "" {
			al.images = append(al.images, Image{title, band, ext})
		}
	})
}

func (image *Image) download (url string) {
	isPathOrCreate(url)
	newFile, _ := os.Create(url + "/" + image.name + image.ext)
	resp, err := http.Get(image.url)
	if err != nil {
		log.Fatal(1111, err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(newFile, resp.Body)
	if err != nil {
		log.Fatal(2222, err)
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