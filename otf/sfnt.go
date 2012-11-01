package otf

import (
	"bytes"
	"encoding/binary"
	"io"
	"sort"
)

func writeSfntHeader(f *SFNT, w io.WriterAt, offset int64) (ULONG, int64, error) {
	numTables := len(f.Table)
	f.Header.SetNumTables(USHORT(numTables))
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, f.Header)
	bytes := buf.Bytes()
	if _, err := w.WriteAt(bytes, offset); err != nil {
		return 0, 0, err
	}
	return calcCheckSum(bytes), int64(len(bytes)), nil
}

func (f *SFNT) Generate(w io.WriterAt) error {
	total, dictOffset, err := writeSfntHeader(f, w, 0)
	if err != nil {
		return err
	}
	// write tables
	numTables := len(f.Table)
	entry := make([]OffsetEntry, numTables)
	offset := roundUp(dictOffset + int64(binary.Size(entry)))
	tag := make(sort.StringSlice, 0)
	entryMap := make(map[string]*OffsetEntry)
	head := make([]int64, 0)
	for _, v := range f.Table {
		t := v.Tag()
		ts := string(t[:])
		if ts == "head" {
			h := v.(*Head)
			h.Set()
			head = append(head, offset)
		}
		bytes := v.Bytes()
		checksum := calcCheckSum(bytes)
		length := len(bytes)
		if _, err := w.WriteAt(bytes, offset); err != nil {
			return err
		}
		tag = append(tag, ts)
		entryMap[ts] = &OffsetEntry{
			t,
			checksum,
			ULONG(offset),
			ULONG(length),
		}
		offset = roundUp(offset + int64(length))
		total += checksum
	}
	// write table directory
	offset = dictOffset
	buf := new(bytes.Buffer)
	sort.Sort(tag)
	for i, ts := range tag {
		entry[i] = *(entryMap[ts])
	}
	binary.Write(buf, binary.BigEndian, entry)
	bytes := buf.Bytes()
	if _, err := w.WriteAt(bytes, offset); err != nil {
		return err
	}
	total += calcCheckSum(bytes)
	adjust := make([]byte, 4)
	binary.BigEndian.PutUint32(adjust, uint32(checkSumAdjustmentMagic-total))
	for _, offs := range head {
		w.WriteAt(adjust, offs+8)
	}
	return nil
}

type SFNT struct {
	Header SfntHeader
	Table  map[string]Table
}

const checkSumAdjustmentMagic = 0xB1B0AFBA
