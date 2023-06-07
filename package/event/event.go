package event

import (
	"encoding"
	"fmt"
	"time"
)

type Type string

const (
	InvestimentoType Type = "InvestimentoType"
	InvestidorType   Type = "InvestidorType"
)

type Base struct {
	ID       string
	Type     Type
	DateTime time.Time
	Retry    bool
}

// Event ...
type Event interface {
	GetID() string
	GetType() Type
	GetDateTime() time.Time
	SetID(id string)
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

func New(t Type) (Event, error) {
	b := &Base{
		Type: t,
	}

	switch t {

	case InvestimentoType:
		return &InvestimentoType{
			Base: b,
		}, nil

	case InvestidorType:
		return &InvestidorType{
			Base: b,
		}, nil

	}

	return nil, fmt.Errorf("type %v not supported", t)
}

func (o *Base) GetID() string {
	return o.ID
}

func (o *Base) SetID(id string) {
	o.ID = id
}

func (o *Base) GetType() Type {
	return o.Type
}

func (o *Base) GetDateTime() time.Time {
	return o.DateTime
}

func (o *Base) String() string {

	return fmt.Sprintf("id:%s type:%s", o.ID, o.Type)
}
