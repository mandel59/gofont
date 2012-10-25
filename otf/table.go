package otf

import (
	"bytes"
	"encoding/binary"
	"io"
)

func calcCheckSum(t []byte) (s ULONG) {
	for i := 0; i < len(t); i += 4 {
		if i+4 > len(t) {
			var u [4]byte
			copy(u[0:len(t)-i], t[i:len(t)])
			s += ULONG(binary.BigEndian.Uint32(u[:]))
		} else {
			s += ULONG(binary.BigEndian.Uint32(t[i : i+4]))
		}
	}
	return
}

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
	return Table(&UnknownTable{t, b})
}

type Table interface {
	Tag() TAG
	CheckSum() ULONG
	Len() ULONG
	WriteTo(w io.Writer) (n int, err error)
}

type Subtable interface {
	Len() ULONG
	WriteTo(w io.Writer) (n int, err error)
}
