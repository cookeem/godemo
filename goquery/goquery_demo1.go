package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"strings"
	"time"
)

func parseAbsUrl(strRootURL, strRelaURL string) (link string, err error) {
	urlRoot, err := url.Parse(strRootURL)
	if err != nil {
		return
	}
	if strings.HasPrefix(strRelaURL, "/") {
		link = fmt.Sprintf("%s://%s%s", urlRoot.Scheme, urlRoot.Host, strRelaURL)
	} else if !strings.Contains(strRelaURL, "://") {
		arr := strings.Split(strRootURL, "/")
		arr[len(arr)-1] = strRelaURL
		link = strings.Join(arr, "/")
	} else {
		_, err = url.Parse(strRelaURL)
		if err != nil {
			return
		}
		link = strRelaURL
	}
	return link, err
}

func parseListPage(strURL string) (links []string) {
	//strURL := "http://git.oschina.net/cookeem/CookIM/stargazers"
	doc, err := goquery.NewDocument(strURL)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find("div[class='item user-list-item']").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		name := s.Find("div.header").Text()
		name = strings.Replace(name, "\n", "", -1)
		avatar := s.Find("img.avater").AttrOr("src", "")
		link := s.Find("div.header > a").AttrOr("href", "")
		link, _ = parseAbsUrl(strURL, link)
		if link != "" {
			links = append(links, link)
		}
		fmt.Printf("Name %d: %s, %v, %v\n", i, name, avatar, link)
	})
	return links
}

func parseContentPage(strURL string) {
	//strURL := "http://git.oschina.net/fulus"
	doc, err := goquery.NewDocument(strURL)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	name := doc.Find("div.user-info").Text()
	name = strings.Replace(name, "\n", " ", -1)
	numText := doc.Find("div[class='git-user-infodata'] > div[class='ui grid'] > div.four").Text()
	numText = strings.Replace(numText, "\n", " ", -1)
	fmt.Println("## Name:", name, "## Tags:", numText)
	fmt.Println("## Link:", strURL, time.Now())
}

func main() {
	strURL := "http://git.oschina.net/cookeem/CookIM/stargazers?page=4"
	links := parseListPage(strURL)
	fmt.Println("##########################")
	fmt.Println("##########################")

	//jobs为带缓冲channel
	jobs := make(chan string, 10)
	numOfWorkers := 10
	for w := 1; w <= numOfWorkers; w++ {
		//启动numOfWorkers个goroutine
		go func() {
			for link := range jobs {
				parseContentPage(link)
			}
		}()
	}

	//把job分配给goroutine
	for _, link := range links {
		jobs <- link
	}

	//在jobs写入的程序段进行channel关闭
	close(jobs)
}
