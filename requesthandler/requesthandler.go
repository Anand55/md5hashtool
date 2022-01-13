package requesthandler

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func MakeHttpRequest(wg *sync.WaitGroup, urlChan <-chan string) error {
	defer wg.Done()
	for url := range urlChan {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error making http request: ", err)
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reaading response body: ", err)
			return err
		}
		responseBody := string(body)

		hash := fmt.Sprintf("%x", md5.Sum([]byte(responseBody)))
		fmt.Println(url + " " + hash)
	}
	return nil
}
