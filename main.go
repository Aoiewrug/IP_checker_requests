package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Aoiewrug/IP_checker_requests/channnales"
	"github.com/Aoiewrug/IP_checker_requests/models"
	"github.com/Aoiewrug/IP_checker_requests/queue"
	"github.com/Aoiewrug/IP_checker_requests/readFile"
	writefile "github.com/Aoiewrug/IP_checker_requests/writeFile"
)

const ThreadLimit = 50

func main() {
	start := time.Now()

	// Credentials are here!
	models.InitializeConfig()

	fmt.Println("Start checking: ", models.Config.Link)

	err := writefile.String(models.Config.Link)
	if err != nil {
		log.Fatalf("Error saving buffer to file: %v", err)
	}

	// Start the saver goroutine
	go queue.Appender()

	// Start worker goroutines
	for i := 0; i < ThreadLimit; i++ {
		go queue.Worker()
	}

	// Extract proxies
	ips, err := readFile.Proxies()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send proxies to IPworkerQ
	go func() {
		for _, ip := range ips {
			time.Sleep(10 * time.Millisecond)
			channnales.IPworkerQChan <- ip
		}
		close(channnales.IPworkerQChan) // Close the channel when done sending
	}()

	// Wait for workers to finish processing
	go func() {
		for i := 0; i < ThreadLimit; i++ {
			<-channnales.IPworkerSigChan
		}
		close(channnales.AppendQChan) // Close SaveQ when all workers are done
	}()

	// Wait for the Append to finish
	<-channnales.AppendSigChan

	//fmt.Println("main buffer", channnales.BufferGlobal)

	// Save the data to a file
	err = writefile.Array(channnales.BufferGlobal)
	if err != nil {
		log.Fatalf("Error saving buffer to file: %v", err)
	}

	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Time taken: %v\n", duration)
	fmt.Println("All work completed and data saved successfully.")

	// Proceed with additional operations
	// For example: open infoTofind file or any other tasks
}
