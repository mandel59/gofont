package otf

import (
	"encoding/binary"
	"io"
	"bytes"
	"sort"
)

func writePadding(w io.Writer, t Table) (n int, err error) {
	n, err = t.WriteTo(w)
	if err != nil {
		return n, err
	}
	l := ((n + 3) &^ 3) - n
	m, err := w.Write(make([]byte, l))
	n += m
	return n, err
}

func (o OTF) Write(w io.Writer) error {
	o.Setup()
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, o.OffsetTable)
	if err != nil {
		return err
	}
	n, err := w.Write(buf.Bytes())
	offs := n
	if err != nil {
		return err
	}
	for _, rec := range o.TableRecords {
		err = binary.Write(w, binary.BigEndian, *rec)
		if err != nil {
			return err
		}
		offs += 16
	}
	for _, t := range o.Tables {
		buf.Reset()
		_, err := writePadding(buf, t)
		if err != nil {
			return err
		}
		n, err = w.Write(buf.Bytes())
		offs += n
		if err != nil {
			return err
		}
	}
	return err
}

func (o OTF) Setup() {
	o.NumTables = USHORT(len(o.Tables))
	o.OffsetTable.SetupNumTables()
	o.TableRecords = make(sliceTableRecord, len(o.Tables))
	offset := 12 + 16 * ULONG(o.NumTables)
	var head int
	var sum ULONG
	for i, t := range o.Tables {
		tag := t.Tag()
		if tag == TAG_HEAD {
			head = i
			o.Tables[head].(*Head).CheckSumAdjustment = 0
		}
		checkSum := t.CheckSum()
		length := t.Len()
		o.TableRecords[i] = &TableRecord{
			tag,
			checkSum,
			offset,
			length,
		}
		offset += (length + 3) &^ 3
		sum += checkSum
	}
	sort.Sort(byTagSort{o.TableRecords})
	for _, r := range o.TableRecords {
		sum += ULONG(r.Tag[0])<<12 + ULONG(r.Tag[1])<<8
		sum += ULONG(r.Tag[2])<<4 + ULONG(r.Tag[3])
		sum += r.CheckSum + r.Offset + r.Length
	}
	sum += ULONG(o.SfntVersion)
	sum += ULONG(o.NumTables)<<8 + ULONG(o.SearchRange)
	sum += ULONG(o.EntrySelector)<<8 + ULONG(o.RangeShift)
	o.Tables[head].(*Head).CheckSumAdjustment = 0xB1B0AFBA - sum
}

type OTF struct {
	OffsetTable
	TableRecords sliceTableRecord
	Tables       []Table
}
