package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate processing of tickets
func ticketProcessing(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of PersonID %d with total cost %d\n", req.numTickets, req.numTickets, req.cost)
		heavyTask(5)
		results <- req.personID
	}
}

func heavyTask(seconds int) {
	end := time.Now().Add(time.Duration(seconds) * time.Second)
	x := 0.0
	for time.Now().Before(end) {
		for i := 1; i < 1_000_000; i++ {
			x += float64(i) * 1.0000001 / float64(i+1)
		}
	}
	fmt.Println("done", x)
}

func main() {
	numRequests := 15
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	numWorkers := 6

	// start ticket processor/worker
	for range numWorkers {
		go ticketProcessing(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i := range numRequests {
		ticketRequests <- ticketRequest{i + 1, (i + 1) * 2, (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests{
		fmt.Printf("Ticket for PersonID %d processed successfully!\n", <-ticketResults)
	}

}
