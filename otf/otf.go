package otf

import (
	"bytes"
	"encoding/binary"
	"io"
	"sort"
)

func (o OTF) Write(w io.Writer) error {
	o.Setup()
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, o.SfntHeader)
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
		n, err = w.Write(t.Bytes())
		offs += n
		if err != nil {
			return err
		}
	}
	return err
}

func (o OTF) Setup() {
	o.SfntHeader.Set(VERSION_1_0, USHORT(len(o.Tables)))
	o.TableRecords = make(sliceTableRecord, len(o.Tables))
	offset := 12 + 16*ULONG(len(o.Tables))
	var head int
	var sum ULONG
	for i, t := range o.Tables {
		tag := t.Tag()
		if tag == TAG_HEAD {
			head = i
			o.Tables[head].(*Head).CheckSumAdjustment = 0
		}
		b := t.Bytes()
		checkSum := calcCheckSum(b)
		length := len(b)
		o.TableRecords[i] = &OffsetTableEntry{
			tag,
			checkSum,
			offset,
			ULONG(length),
		}
		offset += ULONG(roundUp(length))
		sum += checkSum
	}
	sort.Sort(byTagSort{o.TableRecords})
	for _, r := range o.TableRecords {
		sum += ULONG(r.Tag[0])<<12 + ULONG(r.Tag[1])<<8
		sum += ULONG(r.Tag[2])<<4 + ULONG(r.Tag[3])
		sum += r.CheckSum + r.Offset + r.Length
	}
	sum += ULONG(o.SfntHeader.SfntVersion)
	sum += ULONG(o.SfntHeader.NumTables)<<8 + ULONG(o.SfntHeader.SearchRange)
	sum += ULONG(o.SfntHeader.EntrySelector)<<8 + ULONG(o.SfntHeader.RangeShift)
	o.Tables[head].(*Head).CheckSumAdjustment = checkSumAdjustmentMagic - sum
}

type OTF struct {
	SfntHeader   SfntHeader
	TableRecords sliceTableRecord
	Tables       []Table
}

const checkSumAdjustmentMagic = 0xB1B0AFBA
