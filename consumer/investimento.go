package handler

import (
	"errors"
	"fmt"

	evt "event"
)

type InvestimentoHandler struct {
}

// NewInvestimentoHandler ...
func NewInvestimentoHandler() Handler {
	return &InvestimentoHandler{}
}

func (h *InvestimentoHandler) Handle(e evt.Event, retry bool) error {
	event, ok := e.(*evt.InvestimentoEvent)

	if !ok {
		return fmt.Errorf("incorrect event type")
	}

	if event.UserID == 5 && !retry {
		return errors.New("Falhou")
	}

	fmt.Printf("completed Investimento %+v UserID: %v\n", event, event.UserID)

	return nil
}
