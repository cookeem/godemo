package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"time"
)

func main() {
	t1 := time.Now()


	// pageLoadStrategy:
	// normal - waits for document.readyState to be ‘complete’. This value is used by default.
	// eager - will abort the wait when document.readyState is ‘interactive’ instead of waiting for ‘complete’.
	// none - will abort the wait immediately, without waiting for any of the page to load.
	caps := selenium.Capabilities{
		"browserName":      "firefox", //phantomjs
		"pageLoadStrategy": "eager",
	}
	// docker run -d -p 4444:4444 --name selenium-hub selenium/hub
	// wait 3 seconds
	// docker run -d -p 14444:4444 --link selenium-hub:hub selenium/node-phantomjs:latest
	// docker run -d -p 15555:5555 --link selenium-hub:hub selenium/node-firefox:latest
	//
	// use codes below:
	wd, err := selenium.NewRemote(caps, "http://localhost:15555/wd/hub/")
	//wd, err := selenium.NewRemote(caps, "http://localhost:15555")

	// docker run -d -p 5900:5900 -p 4444:4444 selenium/standalone-firefox-debug:latest
	// open vnc://:secret@localhost:5900/
	//
	// use codes below:
	// wd, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub/")
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}

	// Navigate to the simple playground interface.
	link := "https://www.baidu.com/"
	wd.SetPageLoadTimeout(time.Second * 1)
	if err := wd.Get(link); err != nil {
		panic(err)
	}
	sid1 := wd.SessionID()

	title, err := wd.Title()
	fmt.Println("############## title ################")
	fmt.Println(title)

	source, _ := wd.PageSource()
	fmt.Println("############## source ################")
	fmt.Println(source)

	wd.NewSession()
	sid2 := wd.SessionID()

	link = "https://www.sogou.com/"
	if err := wd.Get(link); err != nil {
		panic(err)
	}

	wd.Status()

	searchBox, err := wd.FindElement(selenium.ByID, "search-box")
	if err != nil {
		panic(err)
	}
	searchBoxHtml, err := searchBox.GetAttribute("outerHTML")
	if err != nil {
		panic(err)
	}
	fmt.Println("############## searchBoxHtml ################")
	fmt.Println(searchBoxHtml)

	wd.SwitchSession(sid1)
	ss, _ := wd.Screenshot()
	ioutil.WriteFile(fmt.Sprintf("%v.png", sid1), ss, 0644)
	wd.Quit()

	wd.SwitchSession(sid2)
	ss, _ = wd.Screenshot()
	ioutil.WriteFile(sid2+".png", ss, 0644)
	wd.Quit()

	fmt.Println("sid1:", sid1)
	fmt.Println("sid2:", sid2)

	//fmt.Println("#############################################################")
	//fmt.Println("#############################################################")
	//fmt.Println("#############################################################")
	//
	//// Get a reference to the text box containing code.
	//elem, err := wd.FindElement(selenium.ByTagName, "html")
	//if err != nil {
	//	panic(err)
	//}
	//
	//elemText, _ := elem.GetAttribute("innerHTML")
	//fmt.Println(elemText)

	//// Remove the boilerplate code already in the text box.
	//if err := elem.Clear(); err != nil {
	//	panic(err)
	//}
	//
	//// Enter some new code in text box.
	//err = elem.SendKeys(`
	//	package main
	//	import "fmt"
	//	func main() {
	//		fmt.Println("Hello WebDriver!\n")
	//	}
	//`)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Click the run button.
	//btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	//if err != nil {
	//	panic(err)
	//}
	//if err := btn.Click(); err != nil {
	//	panic(err)
	//}
	//
	//// Wait for the program to finish running and get the output.
	//outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	//if err != nil {
	//	panic(err)
	//}
	//output := ""
	//for {
	//	output, err = outputDiv.Text()
	//	if err != nil {
	//		panic(err)
	//	}
	//	if output != "Waiting for remote server..." {
	//		break
	//	}
	//	time.Sleep(time.Millisecond * 100)
	//}
	//
	//fmt.Printf("Got: %s\n", output)

	fmt.Println("duration: ", time.Now().Sub(t1))
}
