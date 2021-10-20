package process

import (
	"chandylamport/models"
	"fmt"
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
}

func CreateStateManager(pid int, updateStateChan chan models.ProcessEvent, saveGlobalState chan int, markMessageIn chan models.Message, markMessageOut chan int) *StateManager {

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
	}

	go thisStateManager.UpdateState()
	return &thisStateManager
}

func (sm *StateManager) UpdateState() {
	for {
		select {
		case newEvent := <-sm.UpdateStateChan:
			fmt.Println(newEvent)
			sm.processHistory.CurrentEvent += 1
			sm.processHistory.EventHistory = append(sm.processHistory.EventHistory, newEvent)
		case extraDelay := <-sm.SaveGlobalState:
			fmt.Println("Aquí guarda el historico y comienza a grabar los mensajes")
			// Se debe verificar si está en modo snapshot para guardar los mensajes
			var messagesIn []models.Message
			var messagesOut []models.Message
			currentSnapshot := models.Snapshot{
				EventHistory: sm.processHistory,
				MessagesIn:   messagesIn,
				MessagesOut:  messagesOut,
			}
			sm.Snapshots = append(sm.Snapshots, currentSnapshot)
			sm.takingSnapshot = true
			sm.pendingMarks = 2
			sm.MarkMessageOut <- extraDelay
		case MarkMsg := <-sm.MarkMessageIn:
			fmt.Println(MarkMsg)
			if sm.takingSnapshot == true {
				if sm.pendingMarks > 0 {
					sm.pendingMarks -= 1
				} else {
					sm.takingSnapshot = false
				}
			} else {
				sm.takingSnapshot = true
				sm.pendingMarks = 1
				sm.MarkMessageOut <- 100
			}
		}
	}
}

func (sm *StateManager) TakeSnapshot() {
	fmt.Println("Tomar snapshot")
}
