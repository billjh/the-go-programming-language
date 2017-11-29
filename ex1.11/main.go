// go run main.go

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const reqs = 200

func main() {
	alexa, err := os.Open("top-1m.csv")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	urls := csv.NewReader(alexa)
	start := time.Now()
	ch := make(chan string)
	for i := 0; i < reqs; i++ {
		url, err := urls.Read()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		go fetch("http://"+url[1], ch) // start a goroutine
	}
	for i := 0; i < reqs; i++ {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
