package otf

import (
	"encoding/binary"
	"io"
)

func NewTable(t TAG, r *io.SectionReader) Table {
	switch t {
	case TAG_HEAD:
		head := new(Head)
		r.Seek(0, 0)
		err := binary.Read(r, binary.BigEndian, head)
		if err != nil {
			break
		}
		return Table(head)
	}
	return Table(&DefaultTable{t, r})
}

type Table interface {
	Subtable
	Tag() TAG
}

type Subtable interface {
	Bytes() []byte
}
