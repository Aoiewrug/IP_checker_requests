package queue

import (
	"github.com/Aoiewrug/IP_checker_requests/channnales"
	"github.com/Aoiewrug/IP_checker_requests/check"
	"github.com/Aoiewrug/IP_checker_requests/models"
)

func Worker() {
	for {
		select {
		case ip, ok := <-channnales.IPworkerQChan:
			if !ok {
				// If IPworkerQ is closed, exit the goroutine
				channnales.IPworkerSigChan <- struct{}{}
				return
			}
			// Process the work item
			//fmt.Println("Processing:", ip)

			config := models.Config
			config.IP = ip // Assign the IP from the channel to the config
			//fmt.Println("config to process is:", config)

			check.RequestChecker(config)

		}
	}
}

func Appender() {
	for {
		select {
		case ip, ok := <-channnales.AppendQChan:
			if !ok {
				// If SaveQ is closed, exit the goroutine
				channnales.AppendSigChan <- struct{}{}
				return
			}
			// Save the data
			channnales.Mu.Lock()
			//fmt.Println("Appending this IP:", ip)
			channnales.BufferGlobal = append(channnales.BufferGlobal, ip)
			channnales.Mu.Unlock()
			// Simulate saving work
			//time.Sleep(2 * time.Second) // Simulate some work being done
		}
	}
}
