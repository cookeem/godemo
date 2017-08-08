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
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strconv"
)

type FieldType struct {
	Name string
	Pattern string
	IsURL bool
	Attr string
	ReplaceSource string
	ReplaceTarget string
}

type ListType struct {
	Pattern string
	Fields [] FieldType
}

type ContentType struct {
	Fields [] FieldType
}

type ParserConfigType struct {
	List ListType
	Content ContentType
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

func parseListPage(strURL string, listConfig ListType, logFile, logStdout *log.Logger) (items []map[string] string, err error) {
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
		item := make(map[string] string)
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
		logFile.Println("#########################")
		logStdout.Println("#########################")
	})
	return
}

func parseContentPage(strURL string, contentConfig ContentType) (item map[string] string, err error) {
	//strURL := "http://git.oschina.net/fulus"
	item = make(map[string] string)

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
		//必须要有results，不然parseContentPage还没有执行完就会退出
		t1 := time.Now()
		item, err := parseContentPage(strURL, contentConfig)
		if err != nil {
			log.Fatal(err)
			return
		}
		t2 := time.Now()
		d := t2.Sub(t1)
		logFile.Println("worker", w, "take", d, "##", item)
		logStdout.Println("worker", w, "take", d, "##", item)
		results <- "ok"
	}
}

func main() {
	strURL := "http://git.oschina.net/cookeem/CookIM/stargazers?page=1"

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

	listConfig := parserConfig.List
	contentConfig := parserConfig.Content

	items, err := parseListPage(strURL, listConfig, logFile, logStdout)

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
	for i := 1; i <= numOfWorkers * 4; i++ {
		//启动numOfWorkers个goroutine
		go parseContentJob(i, jobs, contentConfig, results, logFile, logStdout)
	}

	//把job分配给goroutine
	for _, item := range items {
		jobs <- item["url"]
	}

	for range items {
		select {
		case <-results:
		case <-time.After(time.Second * 2):
			fmt.Println("!!!!!! timeout 2 seconds")
		}
	}

	//在jobs写入的程序段进行channel关闭
	close(jobs)

}
