package process

import "chandylamport/models"

type StateManager struct {
	processId       string
	processHistory  models.ProcessHistory
	globalState     map[string]models.ProcessHistory
	UpdateStateChan chan models.ProcessEvent
	SaveGlobalState chan bool
}

func CreateStateManager(pid string, updateStateChan chan models.ProcessEvent, saveGlobalState chan bool) *StateManager {

	initEvent := models.ProcessEvent{Description: "Init", Data: ""}
	var eventhistory []models.ProcessEvent
	eventhistory = append(eventhistory, initEvent)
	initHistory := models.ProcessHistory{CurrentEvent: 0, EventHistory: eventhistory}
	initGlobal := map[string]models.ProcessHistory{
		pid: initHistory,
	}

	thisStateManager := StateManager{
		processId:       pid,
		processHistory:  initHistory,
		globalState:     initGlobal,
		UpdateStateChan: updateStateChan,
		SaveGlobalState: saveGlobalState,
	}

	return &thisStateManager
}

func (sm *StateManager) UpdateState() {
	for {
		select {
		case newEvent := <-sm.UpdateStateChan:
			sm.processHistory.CurrentEvent += 1
			sm.processHistory.EventHistory = append(sm.processHistory.EventHistory, newEvent)
		case <-sm.SaveGlobalState:
			sm.globalState[sm.processId] = sm.processHistory
			//case save incomming messages
		}
	}
}
