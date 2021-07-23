package workerpool

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL string
	Status int
}

func RunWorkerPool() {
	urls := []string{
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing/",
		"https://example.com",
		"https://google.com",
	}

	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w:=1; w<4; w++{
		go crawl(w, jobs, results)
	}

	for _, url:=range urls{
		jobs <- Site{URL: url}
	}
	close(jobs)
	for w := 1; w<5; w++ {
		result:=<-results
		fmt.Println(result)
	}
	close(results)
}

func crawl(w int, jobs <-chan Site, results chan <- Result) {
	for site:=range jobs{
		log.Printf("Worker Id: %d\n", w)
		resp, err := http.Get(site.URL)
		if err!=nil {
			log.Println(err)
		}
		results<-Result{URL: site.URL, Status: resp.StatusCode}
	}
}
