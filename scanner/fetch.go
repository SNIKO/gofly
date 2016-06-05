package main

import (
	"fmt"
	"github.com/sniko/gofly/agents"
)

var providers = []agents.Agent{
	agents.Momondo{},
	agents.OneTwoTrip{},
}

func fetcher(trips <-chan TripInfo, fares chan<- agents.Fares) {
	defer close(fares)

	for tripInfo := range trips {
		tripFares := searchAll(tripInfo, providers...)
		fmt.Printf("%d fares have been loaded for route '%s'.\n", len(tripFares), tripInfo)
		fares <- tripFares
	}

	fmt.Printf("All fares have been loaded.\n")
}

func searchAll(trip TripInfo, providers... agents.Agent) (result  agents.Fares) {
	c := make(chan agents.Fares, len(providers))

	for i := range providers {
		agent := providers[i]
		go func() {
			c <- search(agent, trip)
		}()
	}

	for i:=0; i<len(providers); i++ {
		fares := <-c
		result = append(result, fares...)
	}

	return
}

func search(agent agents.Agent, trip TripInfo) agents.Fares {
	fmt.Printf("%s: Searching flights for route '%s'\n", agent.Name(), trip)

	rates, err := agent.Search(trip.Route)
	if (err != nil) {
		fmt.Printf("%s: An error occurred when loading fares for route '%s': %s\n", agent.Name(), trip, err)
		return agents.Fares{}
	}

	fmt.Printf("%s: %d fares for route '%s' have been loaded\n", agent.Name(), len(rates), trip)
	return rates
}