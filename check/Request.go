package check

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Aoiewrug/IP_checker_requests/channnales"
	"github.com/Aoiewrug/IP_checker_requests/models"
)

func RequestChecker(info models.ProxyChanStruct) {

	//println("trying to open ", info.Link)

	proxy, err := url.Parse("http://" + info.Creds + "@" + info.IP + ":" + info.Port)
	if err != nil {
		return
	}
	proxyFunc := http.ProxyURL(proxy)
	tr := &http.Transport{
		MaxIdleConns:       1000,
		IdleConnTimeout:    3 * time.Second,
		DisableCompression: true,
		Proxy:              proxyFunc,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
	resp, err := client.Get(info.Link)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	/*
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		fmt.Print(string(bytes))
	*/

	message := fmt.Sprintf("Response: %d %s Proxy: %s", resp.StatusCode, http.StatusText(resp.StatusCode), info.IP)

	//fmt.Println("test", message)

	// Send the message to the Save channel
	channnales.AppendQChan <- message
}
