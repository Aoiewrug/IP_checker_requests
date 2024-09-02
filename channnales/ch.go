package channnales

import (
	"sync"

	"github.com/Aoiewrug/IP_checker_requests/models"
)

var (
	DataChan = make(chan models.ProxyChanStruct)

	IPworkerQChan   = make(chan string)
	IPworkerSigChan = make(chan struct{})

	AppendQChan   = make(chan string) // Append to the golbal buffer
	AppendSigChan = make(chan struct{})

	BufferGlobal []string
	Mu           sync.Mutex // Mutex to protect BufferGlobal
)
