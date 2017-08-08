package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"os"
	"strings"
	"runtime"
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

func parseListPage(strURL string, logFile, logStdout *log.Logger) (links []string, err error) {
	//strURL := "http://git.oschina.net/cookeem/CookIM/stargazers"
	document, err := goquery.NewDocument(strURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Find the review items
	document.Find("div[class='item user-list-item']").Each(func(i int, s *goquery.Selection) {
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
	return
}

func parseContentPage(strURL string) (name, numText string, err error) {
	//strURL := "http://git.oschina.net/fulus"
	document, err := goquery.NewDocument(strURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Find the review items
	name = document.Find("div.user-info").Text()
	name = strings.Replace(name, "\n", " ", -1)
	numText = document.Find("div[class='git-user-infodata'] > div[class='ui grid'] > div.four").Text()
	numText = strings.Replace(numText, "\n", " ", -1)
	return
}

func parseContentJob(w int, jobs chan string, results chan string, logFile, logStdout *log.Logger) {
	for link := range jobs {
		//必须要有results，不然parseContentPage还没有执行完就会退出
		t1 := time.Now()
		name, numText, err := parseContentPage(link)
		if err != nil {
			log.Fatal(err)
			return
		}
		t2 := time.Now()
		d := t2.Sub(t1)
		logFile.Println("worker", w, "take", d,"## Name:", name, "## Tags:", numText)
		logStdout.Println("worker", w, "take", d, "## Name:", name, "## Tags:", numText)
		results <- "ok"
	}
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

	links, err := parseListPage(strURL, logFile, logStdout)

	if err != nil {
		log.Fatal(err)
	}

	//jobs为带缓冲channel
	numOfJobs := len(links)
	numOfWorkers := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfWorkers)
	jobs := make(chan string, numOfJobs)
	results := make(chan string, numOfJobs)
	for i := 1; i <= numOfWorkers * 2; i++ {
		//启动numOfWorkers个goroutine
		go parseContentJob(i, jobs, results, logFile, logStdout)
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
