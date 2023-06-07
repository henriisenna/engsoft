package handler

import (
	"errors"
	"fmt"

	evt "event"
)

type InvestidorHandler struct {
}

// NewInvestidorHandler ...
func NewInvestidorHandler() Handler {
	return &InvestidorHandler{}
}

func (h *InvestidorHandler) Handle(e evt.Event, retry bool) error {
	event, ok := e.(*evt.InvestidorEvent)

	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	if event.UserID == 5 && !retry {
		return errors.New("Falhou")
	}

	fmt.Printf("processed event %+v UserID: %v Investidor:%v \n", event, event.UserID, event.Investidor)

	return nil
}
