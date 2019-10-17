package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	//	"github.com/axgle/mahonia"
)

func main() {
	path := "qindi1.txt"
	baseurl := "https://www.tianxiabachang.cn"
	url := baseurl + "/5_5593/2268166.html"
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		fmt.Println("OpenFile: ", err)
		return
	}

	for i := 0; i < 10000; i++ {
		if !strings.Contains(url, "html") {
			break
		}
		start := time.Now()
		cli := &http.Client{Timeout: time.Duration(10) * time.Second}
		resp := &http.Response{}
		for {
			resp, err = cli.Get(url)
			if err == nil && !resp.Close {
				break
			}
			fmt.Println(err, resp)
		}

		fmt.Printf("%d, %v", i, time.Since(start))
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Println("NewDocumentFromR.eader: ", err)
			return
		}
		resp.Body.Close()
		boxCon := doc.Find(".box_con")

		// get title
		titlestrs := boxCon.Find(".bookname").Text()
		ts := strings.Split(titlestrs, "\n")
		var title, body string
		for _, line := range ts {
			if strings.Contains(line, "第") && strings.Contains(line, "章") {
				title = line[strings.Index(line, "第"):]
				title = title + "\n"
				break
			}
		}

		url = baseurl + boxCon.Find(".bottem1").Find("a:contains(下一章)").Nodes[0].Attr[0].Val

		body = boxCon.Find("#content").Text() + "\n\n"
		//	body = mahonia.NewDecoder("gbk").ConvertString(body)

		f.WriteString(title)
		f.WriteString(body)

		fmt.Printf("  %v", time.Since(start))
		fmt.Println("")
		time.Sleep(time.Second * time.Duration(2))

	}
}
