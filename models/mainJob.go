package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ProcessInfo struct {
	Name string
	Ip   string
	Port string
}

// This element represents the main task of the Process
type MainJob struct {
	ProcessInfo       ProcessInfo   // it's own info
	NetworkInfo       []ProcessInfo // other processes info
	Data              ProcessData
	updateStateChan   chan ProcessEvent
	processMessageOut chan Message
	incomeChan        chan Message
}

// Returns the pointer to a new MainJob struct
func CreateJob(processInfo ProcessInfo, network []ProcessInfo, updateStateChan chan ProcessEvent, processMessageOut chan Message) *MainJob {

	seed := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(seed)
	frecMsg := time.Duration(r.Intn(1000) + 500)
	tickerJob := time.NewTicker(2 * time.Second)                // Increase amount each 2 seconds
	tickerMessage := time.NewTicker(frecMsg * time.Millisecond) // send message with rand time between 0.5 and 1.5 s
	quit := make(chan bool)

	amount := r.Intn(100)
	myJob := MainJob{ProcessInfo: processInfo, NetworkInfo: network, Data: ProcessData{Amount: amount, Mu: sync.Mutex{}}, updateStateChan: updateStateChan}
	go myJob.MockJob(tickerJob, tickerMessage, quit)
	return &myJob
}

func (p *MainJob) MockJob(ticker *time.Ticker, tickerMessage *time.Ticker, quit chan bool) {
	for {
		select {
		// Increase amount
		case <-ticker.C:
			p.Data.Mu.Lock()
			p.Data.Amount += 1
			// update state
			event := ProcessEvent{Description: "Increase amount", Data: fmt.Sprintf("New amount: %v", p.Data.Amount)}
			p.Data.Mu.Unlock()
			p.updateStateChan <- event

		// send message
		case <-tickerMessage.C:
			seed := rand.NewSource(time.Now().UnixMicro())
			r := rand.New(seed)
			amount := r.Intn(10)
			p.Data.Amount -= amount
			receiver := r.Intn(len(p.NetworkInfo))                  // choose receiver randomly
			if p.NetworkInfo[receiver].Port == p.ProcessInfo.Port { // avoid sending message to itself
				if receiver == 0 {
					receiver += 1
				} else {
					receiver -= 1
				}
			}
			msg := NewMessage(p.ProcessInfo.Name, p.NetworkInfo[receiver].Name, amount)
			p.processMessageOut <- msg
			// update state
			event := ProcessEvent{Description: "Amount Sent", Data: fmt.Sprintf("Send : %v to %v", p.Data.Amount, p.NetworkInfo[receiver].Name)}
			p.updateStateChan <- event

		case income := <-p.incomeChan:
			p.Data.Amount += income.Body.(int)
			// update state
			event := ProcessEvent{Description: "Receive Income", Data: fmt.Sprintf("Receive %v from %v", income.Body, income.Sender)}
			p.updateStateChan <- event

		case <-quit:
			ticker.Stop()
			return
		}
	}
}
