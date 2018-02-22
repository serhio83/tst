package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func MakeRequest(url string, ch chan<-string) {
	start := time.Now()
	resp,_ := http.Get(url)
	secs := time.Since(start).Seconds()
	body,_ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.4f elapsed with response length: %d %s status: %d", secs, len(body), url, resp.StatusCode)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _,url := range os.Args[1:]{
		go MakeRequest(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.4f seconds elapsed\n", time.Since(start).Seconds())
}
