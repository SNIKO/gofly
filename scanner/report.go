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
		ReportElasticSearch(config.ElasticSearch.Host, config.ElasticSearch.Port, config.ElasticSearch.Index, fares)
	}

	fmt.Printf("All fares have been reported.\n")
}