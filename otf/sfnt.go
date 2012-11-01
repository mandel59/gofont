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
	return calcCheckSum(bytes), offset + int64(len(bytes)), nil
}

func writeTables(table []Table, w io.WriterAt, offset int64) (map[Table]*OffsetEntry, []int64, ULONG, error) {
	total := ULONG(0)
	entryMap := make(map[Table]*OffsetEntry)
	head := make([]int64, 0)
	for _, v := range table {
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
			return nil, nil, 0, err
		}
		entryMap[v] = &OffsetEntry{
			t,
			checksum,
			ULONG(offset),
			ULONG(length),
		}
		offset = roundUp(offset + int64(length))
		total += checksum
	}
	return entryMap, head, total, nil
}

func checkSumAdjust(w io.WriterAt, head []int64, checksum ULONG) {
	adjust := make([]byte, 4)
	binary.BigEndian.PutUint32(adjust, uint32(checkSumAdjustmentMagic-checksum))
	for _, offs := range head {
		w.WriteAt(adjust, offs+8)
	}
}

func (f *SFNT) Generate(w io.WriterAt) error {
	total, dictOffset, err := writeSfntHeader(f, w, 0)
	if err != nil {
		return err
	}
	numTables := len(f.Table)
	entry := make([]OffsetEntry, numTables)
	offset := roundUp(dictOffset + int64(binary.Size(entry)))
	table := make([]Table, 0)
	for _, v := range f.Table {
		table = append(table, v)
	}
	entryMap, head, total, err := writeTables(table, w, offset)
	if err != nil {
		return err
	}
	// write table directory
	offset = dictOffset
	tag := make(sort.StringSlice, 0)
	for _, v := range f.Table {
		t := v.Tag()
		tag = append(tag, string(t[:]))
	}
	sort.Sort(tag)
	for i, ts := range tag {
		entry[i] = *(entryMap[f.Table[ts]])
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, entry)
	bytes := buf.Bytes()
	if _, err := w.WriteAt(bytes, offset); err != nil {
		return err
	}
	total += calcCheckSum(bytes)
	checkSumAdjust(w, head, total)
	return nil
}

type SFNT struct {
	Header SfntHeader
	Table  map[string]Table
}

const checkSumAdjustmentMagic = 0xB1B0AFBA
