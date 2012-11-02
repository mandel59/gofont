package otf

import (
	"bytes"
	"encoding/binary"
	"io"
	"sort"
	"fmt"
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

func writeTableDirectory(f *SFNT, entryMap map[Table]*OffsetEntry, w io.WriterAt, offset int64) (ULONG, error) {
	tag := make(sort.StringSlice, 0)
	for k, v := range f.Table {
		t := v.Tag()
		ts := string(t[:])
		if k != ts {
			return 0, fmt.Errorf("inconsistent table tag '%s' and '%s'", k, ts)
		}
		tag = append(tag, ts)
	}
	sort.Sort(tag)
	entry := make([]OffsetEntry, len(f.Table))
	for i, ts := range tag {
		entry[i] = *(entryMap[f.Table[ts]])
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, entry)
	bytes := buf.Bytes()
	if _, err := w.WriteAt(bytes, offset); err != nil {
		return 0, err
	}
	return calcCheckSum(bytes), nil
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
	cs, err := writeTableDirectory(f, entryMap, w, dictOffset)
	if err != nil {
		return err
	}
	total += cs
	checkSumAdjust(w, head, total)
	return nil
}

// SFNT is the pair of SfntHeader and map[string]Table.
// The keys of the map are table tag
type SFNT struct {
	Header SfntHeader
	Table  map[string]Table
}

const checkSumAdjustmentMagic = 0xB1B0AFBA
