package process

import (
	"chandylamport/helpers"
	"chandylamport/models"
	"fmt"
	"strings"
)

type StateManager struct {
	processId       int
	processHistory  models.ProcessHistory
	Snapshots       []models.Snapshot
	UpdateStateChan chan models.ProcessEvent
	SaveGlobalState chan int
	MarkMessageIn   chan models.Message
	MarkMessageOut  chan int
	takingSnapshot  bool
	pendingMarks    int
	Logger          *helpers.Logger
}

func CreateStateManager(pid int, updateStateChan chan models.ProcessEvent, saveGlobalState chan int, markMessageIn chan models.Message, markMessageOut chan int, logger *helpers.Logger) *StateManager {

	initEvent := models.ProcessEvent{Description: "Init", Data: ""}
	var eventhistory []models.ProcessEvent
	eventhistory = append(eventhistory, initEvent)
	initHistory := models.ProcessHistory{CurrentEvent: 0, EventHistory: eventhistory}
	var snapshots []models.Snapshot

	thisStateManager := StateManager{
		processId:       pid,
		processHistory:  initHistory,
		Snapshots:       snapshots,
		UpdateStateChan: updateStateChan,
		SaveGlobalState: saveGlobalState,
		MarkMessageIn:   markMessageIn,
		MarkMessageOut:  markMessageOut,
		takingSnapshot:  false,
		pendingMarks:    0,
		Logger:          logger,
	}

	go thisStateManager.UpdateState()
	return &thisStateManager
}

func (sm *StateManager) UpdateState() {
	for {
		select {
		case newEvent := <-sm.UpdateStateChan:
			sm.Logger.Event.Println(newEvent)
			sm.Logger.GoVectLog(fmt.Sprintf("[%v] %v", newEvent.Description, newEvent.Data))
			if !sm.takingSnapshot || newEvent.Description != models.MsgApp {
				sm.processHistory.CurrentEvent += 1
				sm.processHistory.EventHistory = append(sm.processHistory.EventHistory, newEvent)
			} else {
				if strings.Contains(newEvent.Data, "Send") {
					sm.Snapshots[len(sm.Snapshots)-1].MessagesOut = append(sm.Snapshots[0].MessagesOut, newEvent)
				} else {
					sm.Snapshots[len(sm.Snapshots)-1].MessagesIn = append(sm.Snapshots[0].MessagesIn, newEvent)
				}
			}
		case extraDelay := <-sm.SaveGlobalState:
			if !sm.takingSnapshot { // if it's already taking a snapshot just ignore it
				sm.pendingMarks = 2
				sm.TakeSnapshot(extraDelay)
			}
		case MarkMsg := <-sm.MarkMessageIn:
			sm.Logger.Mark.Println(MarkMsg)
			if sm.takingSnapshot {
				sm.pendingMarks -= 1
				if sm.pendingMarks == 0 {
					sm.takingSnapshot = false
					sm.Logger.GoVectLog("Finish taking the snapshot")
					sm.Logger.Snapshot.Println(sm.Snapshots[len(sm.Snapshots)-1]) // quitar
				}
			} else {
				sm.pendingMarks = 1
				sm.TakeSnapshot(100) // Extra delay of 100 ms, just because
			}
		}
	}
}

func (sm *StateManager) TakeSnapshot(extraDelay int) {
	sm.Logger.Mark.Println("Start taking the snapshot")
	sm.Logger.GoVectLog("Start taking the snapshot")
	var messagesIn []models.ProcessEvent
	var messagesOut []models.ProcessEvent
	currentSnapshot := models.Snapshot{
		EventHistory: sm.processHistory,
		MessagesIn:   messagesIn,
		MessagesOut:  messagesOut,
	}
	sm.Snapshots = append(sm.Snapshots, currentSnapshot)
	sm.takingSnapshot = true
	sm.MarkMessageOut <- extraDelay
}
