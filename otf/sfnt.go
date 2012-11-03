package otf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"sort"
)

func writeAt(w io.WriterAt, i interface{}, offset int64) (ULONG, int64, error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, i)
	bytes := buf.Bytes()
	if _, err := w.WriteAt(bytes, offset); err != nil {
		return 0, 0, err
	}
	return calcCheckSum(bytes), offset + int64(len(bytes)), nil
}

func setupTables(f SFNT, table *[]Table, tableSet map[Table]bool) error {
	for k, v := range f {
		if ts := v.Tag().String(); k != ts {
			return fmt.Errorf("inconsistent table tag '%s' and '%s'", k, ts)
		}
		if _, ok := tableSet[v]; ok {
			continue
		}
		tableSet[v] = true
		*table = append(*table, v)
		if err := v.SetUp(f); err != nil {
			return err
		}
	}
	return nil
}

func writeTables(table []Table, w io.WriterAt, offset int64) (map[Table]*OffsetEntry, []int64, ULONG, error) {
	total := ULONG(0)
	entryMap := make(map[Table]*OffsetEntry)
	head := make([]int64, 0)
	for _, v := range table {
		t := v.Tag()
		ts := t.String()
		if ts == "head" {
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

func writeTableDirectory(f SFNT, entryMap map[Table]*OffsetEntry, w io.WriterAt, offset int64) (ULONG, error) {
	tag := make(sort.StringSlice, 0)
	for k, _ := range f {
		tag = append(tag, k)
	}
	sort.Sort(tag)
	entry := make([]OffsetEntry, f.NumTables())
	for i, ts := range tag {
		entry[i] = *(entryMap[f[ts]])
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

func (f SFNT) NumTables() int {
	return len(f)
}

func (f SFNT) WithCFF() bool {
	_, ok := f["CFF "]
	return ok
}

func (f SFNT) Header() *SfntHeader {
	h := new(SfntHeader)
	if f.WithCFF() {
		h.SfntVersion = VERSION_CFF
	} else {
		h.SfntVersion = VERSION_1_0
	}
	h.SetNumTables(USHORT(f.NumTables()))
	return h
}

func (f SFNT) GenerateSFNT(w io.WriterAt, header *SfntHeader) error {
	total, dictOffset, err := writeAt(w, header, 0)
	if err != nil {
		return err
	}
	numTables := f.NumTables()
	entry := make([]OffsetEntry, numTables)
	offset := roundUp(dictOffset + int64(binary.Size(entry)))
	table := make([]Table, 0)
	tableSet := make(map[Table]bool)
	if err := setupTables(f, &table, tableSet); err != nil {
		return err
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

func (f SFNT) Generate(w io.WriterAt) error {
	header := f.Header()
	return f.GenerateSFNT(w, header)
}

// SFNT maps table tag to Table
type SFNT map[string]Table

const checkSumAdjustmentMagic = 0xB1B0AFBA
