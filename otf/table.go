package otf

import (
	"bytes"
	"encoding/binary"
)

func NewTable(t TAG, b []byte) Table {
	switch t {
	case TAG_HEAD:
		head := new(Head)
		buf := bytes.NewBuffer(b)
		err := binary.Read(buf, binary.BigEndian, head)
		if err != nil {
			break
		}
		return Table(head)
	}
	return Table(&DefaultTable{t, b})
}

type Table interface {
	Subtable
	Tag() TAG
}

type Subtable interface {
	Bytes() []byte
}
