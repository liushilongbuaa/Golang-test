package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

func main() {
	path := "fanren.txt"
	baseurl := "http://www.biquwo.org"
	url := baseurl + "/bqw7933/5190239.html"
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		fmt.Println("OpenFile: ", err)
		return
	}

	for i := 0; i < 100000; i++ {
		if !strings.Contains(url, "html") {
			break
		}
		start := time.Now()
		cli := &http.Client{Timeout: time.Duration(10) * time.Second}
		resp := &http.Response{}
		var doc *goquery.Document
		charset := false
		for {
			resp, err = cli.Get(url)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			doc, err = goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				fmt.Println("NewDocumentFromR.eader: ", err)
				continue
			}
			charset = strings.Contains(resp.Header.Get("Content-Type"), "charset=gbk")
			defer resp.Body.Close()
			break
		}
		fmt.Printf("%d, %v", i, time.Since(start))

		boxCon := doc.Find(".box_con")
		var decode func(string) string
		if charset {
			decode = func(a string) string {
				return mahonia.NewDecoder("gbk").ConvertString(a)
			}
		} else {
			decode = func(a string) string {
				return a
			}
		}
		// get title
		title := decode(boxCon.Find(".bookname").Find("h1").Text()) + "\n\n"
		//	xiayizhang := "a:contains(" + mahonia.NewEncoder("bgk").ConvertString("下一章") + ")"
		xiayizhang := "a:contains(下一章)"
		url = baseurl + boxCon.Find(".bottem1").Find(xiayizhang).Nodes[0].Attr[0].Val

		body := decode(boxCon.Find("#content").Text()) + "\n\n"
		title = strings.Replace(title, "聽聽聽聽", "", -1)
		body = strings.Replace(body, "聽聽聽聽", "", -1)
		f.WriteString(title)
		f.WriteString(body)

		fmt.Printf("  %v", time.Since(start))
		fmt.Println("")
		time.Sleep(time.Second * time.Duration(2))
	}
}
