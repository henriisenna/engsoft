package event

import "github.com/vmihailenco/msgpack/v4"

type InvestidorType struct {
	*Base
	UserID  uint64
	Comment string
}

func (o *InvestidorType) MarshalBinary() (data []byte, err error) {
	return msgpack.Marshal(o)
}

func (o *InvestidorType) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, o)
}
