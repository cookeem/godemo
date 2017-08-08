package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Project struct {
	URL       string
	FullURL   string
	Title     string
	UserName  string
	LoginName string
	Lang      string
	Desc      string
	Watches   int
	Stars     int
	Forks     int
}

type User struct {
	URL        string
	FullURL    string
	UserName   string
	LoginName  string
	Followers  int
	Followings int
	Stars      int
	Watches    int
}

type ProjectStar struct {
	ProjectURL string
	LoginName  string
}

func getAbsUrl(strRootURL, strRelaURL string) (link string, err error) {
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

func parseProjectList(strURL string, logStdout *log.Logger) (projects []Project) {
	//strURL := "http://git.oschina.net/explore/recommend?page=1"
	doc, err := goquery.NewDocument(strURL)
	errmsg := ""
	if err != nil {
		errmsg = "parse project list error"
		logStdout.SetPrefix("[ERROR]")
		logStdout.Println(errmsg, err)
		return
	}
	doc.Find("div#git-discover-list > div.item").Each(func(i int, s *goquery.Selection) {
		project := Project{}
		project.URL = s.Find("div.project-title > a").AttrOr("href", "")
		if project.URL != "" {
			project.FullURL, _ = getAbsUrl(strURL, project.URL)
		}
		strUserNameTitle := s.Find("div.project-title > a").Text()
		arrUserNameTitle := strings.Split(strUserNameTitle, "/")
		if len(arrUserNameTitle) > 1 {
			project.UserName = arrUserNameTitle[0]
			project.Title = strings.Join(arrUserNameTitle[1:], "/")
		}
		project.Desc = s.Find("div.project-desc").Text()
		arrURL := strings.Split(project.URL, "/")
		if len(arrURL) == 3 {
			project.LoginName = arrURL[1]
		}
		if project.URL != "" {
			projects = append(projects, project)
		}
		logStdout.SetPrefix("[DEBUG]")
		logStdout.Printf("%d, %#v\n", i, project)
	})
	return
}

func parseProjectContent(project *Project, logStdout *log.Logger) (projectStars []ProjectStar) {
	//strURL := "http://git.oschina.net/moliysdev/MLDPhotoManager"
	doc, err := goquery.NewDocument(project.FullURL)
	errmsg := ""
	if err != nil {
		errmsg = "parse project content error"
		logStdout.SetPrefix("[ERROR]")
		logStdout.Println(errmsg, err)
		return
	}
	project.Watches, _ = strconv.Atoi(doc.Find("span.watch-container a.social-count").Text())
	project.Stars, _ = strconv.Atoi(doc.Find("span.star-container a.social-count").Text())
	project.Forks, _ = strconv.Atoi(doc.Find("span.fork-container a.social-count").Text())

	doc, err = goquery.NewDocument(project.FullURL + "/stargazers")
	errmsg = ""
	if err != nil {
		errmsg = "parse project stargazers error"
		logStdout.SetPrefix("[ERROR]")
		logStdout.Println(errmsg, err)
		return
	}

	doc.Find("div#git-discover-list > div.item").Each(func(i int, s *goquery.Selection) {
		projectStar := ProjectStar{}

		projectStar.ProjectURL = project.URL
		//projectStar.LoginName
		project.URL = s.Find("div.project-title > a").AttrOr("href", "")
		strUserNameTitle := s.Find("div.project-title > a").Text()
		arrUserNameTitle := strings.Split(strUserNameTitle, "/")
		if len(arrUserNameTitle) > 1 {
			project.UserName = arrUserNameTitle[0]
			project.Title = strings.Join(arrUserNameTitle[1:], "/")
		}
		project.Desc = s.Find("div.project-desc").Text()
		arrURL := strings.Split(project.URL, "/")
		if len(arrURL) == 3 {
			project.LoginName = arrURL[1]
		}
		if project.URL != "" {
			projectStars = append(projectStars, projectStar)
		}
		logStdout.SetPrefix("[DEBUG]")
		logStdout.Printf("%d, %#v\n", i, project)
	})
	return
}

func main() {
	logStdOut := log.New(os.Stdout, "[DEBUG]", log.LstdFlags)
	parseProjectList("http://git.oschina.net/explore/recommend?page=1", logStdOut)
}
