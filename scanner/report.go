package main

import (
	"fmt"
	"sync"
	"github.com/sniko/gofly/agents"
)

func reporter(fares <-chan agents.Fares, config Config, wg *sync.WaitGroup) {
	defer wg.Done()
	for fares := range fares {
		fmt.Printf("Reporting %d fares\n", len(fares))
		ReportElasticsearch(config.ElasticSearch.Host, config.ElasticSearch.Port, fares)
	}

	fmt.Println("All fares have been reported.")
}