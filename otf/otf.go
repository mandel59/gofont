package otf

import (
	"encoding/binary"
	"io"
)

func writePadding(w io.Writer, t Table) (n int, err error) {
	n, err = t.WriteTo(w)
	if err != nil {
		return n, err
	}
	l := ((n+3)&^3)-n
	m, err := w.Write(make([]byte, l))
	n += m
	return n, err
}

func (o OTF) WriteTo(w io.Writer) (n int, err error) {
	binary.Write(w, binary.BigEndian, o.OffsetTable)
	for _, t := range o.Tables {
		m, err := writePadding(w, t)
		n += m
		if err != nil {
			return n, err
		}
	}
	return n, err
}

func (o OTF) Setup() {
	o.OffsetTable = NewOffsetTable(o.SfntVersion, USHORT(len(o.Tables)))
	o.TableRecords = make([]TableRecord, len(o.Tables))
	offset := ULONG(binary.Size(o.OffsetTable))+3&^3
	offset += ULONG(binary.Size(o.TableRecords))+3&^3
	var head int
	for i, t := range o.Tables {
		tag := t.Tag()
		if tag == TAG_HEAD {
			head = i
			o.Tables[head].(*Head).CheckSumAdjustment = 0
		}
		checkSum := t.CheckSum()
		length := t.Len()
		o.TableRecords[i] = TableRecord{
			tag,
			checkSum,
			offset,
			length,
		}
		offset += (length+3)&^3
	}
	for _, t := range o.Tables {
		s := t.CheckSum()
	}
	o.Tables[head].(*Head).CheckSumAdjustment = 0
}

type OTF struct {
	OffsetTable
	TableRecords []TableRecord
	Tables []Table
}
