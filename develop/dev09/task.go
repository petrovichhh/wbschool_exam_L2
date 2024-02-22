package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Downloader struct {
	urls []string
}

func (d *Downloader) readUrl() error {
	args := os.Args

	if len(args) == 1 {
		return errors.New("missing URL")
	}

	for _, url := range args[1:] {
		d.urls = append(d.urls, url)
	}

	return nil
}

func extractLinks(body io.Reader) ([]string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}

func clientGet(url string, filename string) {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	links, err := extractLinks(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	for _, link := range links {
		clientGet(link, filename)
	}

	if strings.HasSuffix(url, "/") {
		filename = "index"
	}

	file, err := os.OpenFile(filename+".html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	var d Downloader

	err := d.readUrl()
	if err != nil {
		log.Println("wget: missing URL\nUsage: wget [OPTION]... [URL]...\n\nTry `wget --help' for more options.")
	} else {
		for _, url := range d.urls {
			filename := strings.Map(func(r rune) rune {
				if r == '/' || r == ':' {
					return -1
				}
				return r
			}, url)
			clientGet(url, filename)
		}
	}
}
