package models

import (
	"fmt"
	"sync"

	"github.com/Aoiewrug/IP_checker_requests/readFile"
)

// Global variable to hold the config
var Config ProxyChanStruct

// Initialize the config variable with default values
var once sync.Once

func InitializeConfig() {
	once.Do(func() {
		// Extract links
		link, err := readFile.Links()
		if err != nil {
			fmt.Println(err)
			return
		}

		Config = ProxyChanStruct{
			Port:  "5432",
			Creds: "ikolomeytsev:q0kkb04j",
			Link:  link,
		}

	})
}

type ProxyChanStruct struct {
	Request_Browser string
	// 	Request = request code,
	//	BrowserCode = browser code,
	//	BrowserElem = browser find element on a page
	//	BrowserText = browser find text on a page

	Port  string
	IP    string
	Creds string

	Link    string
	Element string
	Text    string
}
