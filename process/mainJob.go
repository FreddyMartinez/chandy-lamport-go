package process

import (
	"chandylamport/models"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// This element represents the main task of the Process
type MainJob struct {
	ProcessInfo       models.ProcessInfo   // it's own info
	NetworkInfo       []models.ProcessInfo // other processes info
	Data              models.ProcessData
	updateStateChan   chan models.ProcessEvent
	processMessageOut chan models.Message
	processMessageIn  chan models.Message
	markMessageOut    chan int
	ticker            *time.Ticker
}

// Returns the pointer to a new MainJob struct
func CreateJob(processInfo models.ProcessInfo, network []models.ProcessInfo, updateStateChan chan models.ProcessEvent, processMessageIn chan models.Message, processMessageOut chan models.Message, taskList []models.Task, markMessageOut chan int, quit chan bool) *MainJob {
	seed := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(seed)
	tickerJob := time.NewTicker(2 * time.Second) // Increase amount each 2 seconds

	amount := r.Intn(100)
	myJob := MainJob{
		ProcessInfo:       processInfo,
		NetworkInfo:       network,
		Data:              models.ProcessData{Amount: amount, Mu: sync.Mutex{}},
		updateStateChan:   updateStateChan,
		processMessageOut: processMessageOut,
		processMessageIn:  processMessageIn,
		markMessageOut:    markMessageOut,
		ticker:            tickerJob,
	}
	go myJob.MockJob()
	go myJob.ExecuteTaskList(taskList, quit)
	return &myJob
}

func (p *MainJob) MockJob() {
	for {
		select {
		// Increase amount
		case <-p.ticker.C:
			p.Data.Mu.Lock()
			p.Data.Amount += 1
			// update state
			event := models.ProcessEvent{Description: "Increase one.", Data: fmt.Sprintf("New amount: %v", p.Data.Amount)}
			p.Data.Mu.Unlock()
			p.updateStateChan <- event

		case income := <-p.processMessageIn:
			p.Data.Mu.Lock()
			p.Data.Amount += income.Body
			p.Data.Mu.Unlock()
			// update state
			event := models.ProcessEvent{Description: models.MsgApp, Data: fmt.Sprintf("Receive %v from %v, balance %v", income.Body, income.Sender, p.Data.Amount)}
			p.updateStateChan <- event

		}
	}
}

func (p *MainJob) ExecuteTaskList(taskList []models.Task, quit chan bool) {
	for _, task := range taskList {
		time.Sleep(time.Duration(task.Time) * time.Second)
		// if message is App type send amount
		if task.MsgType == models.MsgApp {
			p.Data.Mu.Lock()
			p.Data.Amount -= task.Value
			msg := models.NewAppMessage(p.ProcessInfo.Name, p.NetworkInfo[task.Receiver].Name, task.Value, task.NetworkDelay)
			p.processMessageOut <- msg
			// update state
			event := models.ProcessEvent{Description: models.MsgApp, Data: fmt.Sprintf("Send %v to %v, balance %v", task.Value, p.NetworkInfo[task.Receiver].Name, p.Data.Amount)}
			p.updateStateChan <- event
			p.Data.Mu.Unlock()
		} else if task.MsgType == models.MsgMark {
			p.markMessageOut <- task.NetworkDelay
		}
	}

	time.Sleep(5 * time.Second)
	p.ticker.Stop()
	fmt.Println("Bye")
	quit <- true
}
