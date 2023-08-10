package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ReceiptMatch struct{}

type ReceiptMatchState struct {
	chanID         string
	receiptsToSend receiptChan
}

func (m *ReceiptMatch) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	state := &ReceiptMatchState{
		chanID:         "singleton-match",
		receiptsToSend: make(receiptChan, 100),
	}
	globalReceiptsDispatcher.subscribe(state.chanID, state.receiptsToSend)
	tickRate := 1 // 1 tick per second = 1 MatchLoop func invocations per second
	label := ""
	return state, tickRate, label
}

func (m *ReceiptMatch) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, stateIface interface{}, presences []runtime.Presence) interface{} {
	state, ok := stateIface.(*ReceiptMatchState)
	if !ok {
		logger.Error("state not a valid lobby state object")
		return nil
	}

	return state
}

func (m *ReceiptMatch) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, stateIface interface{}, presences []runtime.Presence) interface{} {
	state, ok := stateIface.(*ReceiptMatchState)
	if !ok {
		logger.Error("state not a valid lobby state object")
		return nil
	}

	return state
}

func (m *ReceiptMatch) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, stateIface interface{}, messages []runtime.MatchData) interface{} {
	state, ok := stateIface.(*ReceiptMatchState)
	if !ok {
		logger.Error("state not a valid lobby state object")
		return nil
	}
	var receiptsToSend []*Receipt

	more := true
	for more {
		select {
		case r := <-state.receiptsToSend:
			receiptsToSend = append(receiptsToSend, r)
		default:
			more = false
		}
	}

	for _, r := range receiptsToSend {
		buf, err := json.Marshal(r)
		if err != nil {
			continue
		}
		dispatcher.BroadcastMessage(101, buf, nil, nil, true)
	}

	return state
}
func (m *ReceiptMatch) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	return state, true, ""
}

func (m *ReceiptMatch) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	return nil
}

func (m *ReceiptMatch) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return nil, ""
}
