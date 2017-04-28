package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"os"
	"strings"
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

func parseListPage(strURL string, logFile, logStdout *log.Logger) (links []string) {
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
		logFile.Printf("Name %d: %s, %v, %v\n", i, name, avatar, link)
		logStdout.Printf("Name %d: %s, %v, %v\n", i, name, avatar, link)
	})
	return links
}

func parseContentPage(strURL string, results chan string, logFile, logStdout *log.Logger) {
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
	logFile.Println("## Name:", name, "## Tags:", numText)
	logStdout.Println("## Name:", name, "## Tags:", numText)
	results <- strURL
}

func main() {
	strURL := "http://git.oschina.net/cookeem/CookIM/stargazers?page=1"

	fileName := "goquery.log"
	file, err := os.Create(fileName)
	if err != nil {
		log.SetPrefix("[ERROR]")
		log.SetFlags(log.LstdFlags)
		log.Println(err)
	}
	defer file.Close()
	logFile := log.New(file, "[DEBUG]", log.LstdFlags)
	logStdout := log.New(os.Stdout, "[DEBUG]", log.LstdFlags)

	links := parseListPage(strURL, logFile, logStdout)

	//jobs为带缓冲channel
	numOfJobs := len(links)
	numOfWorkers := 5
	jobs := make(chan string, numOfJobs)
	results := make(chan string, numOfJobs)
	for i := 1; i <= numOfWorkers; i++ {
		//启动numOfWorkers个goroutine
		go func() {
			for link := range jobs {
				//必须要有results，不然parseContentPage还没有执行完就会退出
				parseContentPage(link, results, logFile, logStdout)
			}
		}()
	}

	//把job分配给goroutine
	for _, link := range links {
		jobs <- link
	}
	for a := 1; a <= numOfJobs; a++ {
		<-results
	}

	//在jobs写入的程序段进行channel关闭
	close(jobs)

}
