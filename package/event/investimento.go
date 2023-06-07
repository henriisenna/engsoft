package event

import "github.com/vmihailenco/msgpack/v4"

type InvestimentoType struct {
	*Base
	UserID uint64
}

func (o *InvestimentoType) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

func (o *InvestimentoType) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
