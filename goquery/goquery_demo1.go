package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type FieldType struct {
	Name          string
	Pattern       string
	IsURL         bool
	Attr          string
	ReplaceSource string
	ReplaceTarget string
}

type ListPagesType struct {
	Pattern string
	IsURL   bool
	Attr    string
	Prefix  string
}

type ListType struct {
	Pattern string
	Fields  []FieldType
}

type ContentType struct {
	Fields []FieldType
}

type ParserConfigType struct {
	URL             string
	Timeout         int
	ContentUrlField string
	ListPages       ListPagesType
	List            ListType
	Content         ContentType
}

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

func parseListPage(strURL string, listConfig ListType, listPagesConfig ListPagesType, logFile, logStdout *log.Logger) (items []map[string]string, pages []string, err error) {
	document, err := goquery.NewDocument(strURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	if listConfig.Pattern == "" || len(listConfig.Fields) == 0 {
		log.Fatal("crawel.yaml config error")
	}

	document.Find(listConfig.Pattern).Each(func(i int, s *goquery.Selection) {
		logStr := strconv.Itoa(i) + " "
		item := make(map[string]string)
		for _, field := range listConfig.Fields {
			k := field.Name
			v := ""
			node := s.Find(field.Pattern)
			if field.Attr == "" {
				v = node.Text()
			} else {
				v = node.AttrOr(field.Attr, "")
			}
			if field.ReplaceSource != "" {
				v = strings.Replace(v, field.ReplaceSource, field.ReplaceTarget, -1)
			}
			if field.IsURL && v != "" {
				v, _ = parseAbsUrl(strURL, v)
			}
			item[k] = v
			logStr = logStr + k + ": " + v + ", "
		}

		items = append(items, item)
		logFile.Println(logStr)
		logStdout.Println(logStr)
	})

	document.Find(listPagesConfig.Pattern).Each(func(i int, s *goquery.Selection) {
		link := s.AttrOr(listPagesConfig.Attr, "")
		if strings.Index(link, listPagesConfig.Prefix) == 0 {
			link, _ = parseAbsUrl(strURL, link)
			pages = append(pages, link)
		}
	})
	return
}

func parseContentPage(strURL string, contentConfig ContentType) (item map[string]string, err error) {
	//strURL := "http://git.oschina.net/fulus"
	item = make(map[string]string)

	document, err := goquery.NewDocument(strURL)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, field := range contentConfig.Fields {
		k := field.Name
		v := ""
		node := document.Find(field.Pattern)
		if field.Attr == "" {
			v = node.Text()
		} else {
			v = node.AttrOr(field.Attr, "")
		}
		if field.ReplaceSource != "" {
			v = strings.Replace(v, field.ReplaceSource, field.ReplaceTarget, -1)
		}
		item[k] = v
	}
	return
}

func parseContentJob(w int, jobs chan string, contentConfig ContentType, results chan string, logFile, logStdout *log.Logger) {
	for strURL := range jobs {
		t1 := time.Now()
		item, err := parseContentPage(strURL, contentConfig)
		if err != nil {
			log.Fatal(err)
			return
		}
		t2 := time.Now()
		duration := t2.Sub(t1)
		logFile.Println("worker", w, "take", duration, "##", strURL, item)
		logStdout.Println("worker", w, "take", duration, "##", strURL, item)
		results <- "ok"
	}
}

func main() {
	t1 := time.Now()
	logName := "goquery.log"
	fileLog, err := os.Create(logName)
	if err != nil {
		log.SetPrefix("[ERROR]")
		log.SetFlags(log.LstdFlags)
		log.Println(err)
	}
	defer fileLog.Close()
	logFile := log.New(fileLog, "[DEBUG]", log.LstdFlags)
	logStdout := log.New(os.Stdout, "[DEBUG]", log.LstdFlags)

	fileYaml, err := os.Open("conf/crawel.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer fileYaml.Close()

	byteYaml, err := ioutil.ReadAll(fileYaml)
	parserConfig := ParserConfigType{}
	err = yaml.Unmarshal(byteYaml, &parserConfig)
	if err != nil {
		log.Fatal(err)
	}

	strURL := parserConfig.URL
	timeout := parserConfig.Timeout
	contentUrlField := parserConfig.ContentUrlField
	listPagesConfig := parserConfig.ListPages
	listConfig := parserConfig.List
	contentConfig := parserConfig.Content

	items, pages, err := parseListPage(strURL, listConfig, listPagesConfig, logFile, logStdout)
	for _, page := range pages {
		fmt.Println(page)
	}

	if err != nil {
		log.Fatal(err)
	}

	logFile.Println("####################################")
	logStdout.Println("####################################")

	//jobs为带缓冲channel
	numOfJobs := len(items)
	numOfWorkers := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfWorkers)
	jobs := make(chan string, numOfJobs)
	results := make(chan string, numOfJobs)
	for i := 1; i <= numOfWorkers*4; i++ {
		//启动numOfWorkers个goroutine
		go parseContentJob(i, jobs, contentConfig, results, logFile, logStdout)
	}

	//把job分配给goroutine
	for _, item := range items {
		jobs <- item[contentUrlField]
	}

	for range items {
		select {
		case <-results:
		case <-time.After(time.Second * time.Duration(timeout)):
			fmt.Println("!!!!!! timeout", timeout, "seconds")
		}
	}

	//在jobs写入的程序段进行channel关闭
	close(jobs)

	t2 := time.Now()
	duration := t2.Sub(t1)
	logFile.Println("total duration:", duration)
	logStdout.Println("total duration:", duration)

}
