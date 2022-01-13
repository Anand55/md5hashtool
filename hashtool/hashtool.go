package hashtool

import (
	"flag"
	"sync"

	"github.com/Anand55/md5hashtool/requesthandler"
)

var wg sync.WaitGroup

func Start() {
	var parallelRequestCount int
	flag.IntVar(&parallelRequestCount, "parallel", 10, "Parallel requests to be executed")
	flag.Parse()
	args := flag.Args()
	request := make(chan string, len(args))

	if parallelRequestCount > len(args) {
		parallelRequestCount = len(args)
	}
	for _, url := range args {
		request <- url
	}
	close(request)

	for p := 0; p < parallelRequestCount; p++ {
		wg.Add(1)
		go requesthandler.MakeHttpRequest(&wg, request)
	}
	wg.Wait()
}
