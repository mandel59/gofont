package otf

import (
	"bytes"
	"encoding/binary"
	"io"
	"sort"
)

type TTCHeader struct {
	TTCTag   TAG
	Version  FIXED
	NumFonts ULONG
}

type TTCOffsetTable []ULONG

type TTCHeaderDsig struct {
	UIDsigTag    ULONG
	UIDsigLength ULONG
	UIDsigOffset ULONG
}

var TAG_TTC TAG = TAG{'t', 't', 'c', 'f'}

func writeTTCHeader(o OTF, w io.WriterAt) (ULONG, int64, error) {
	numFonts := len(o)
	header := TTCHeader{
		TAG_TTC,
		VERSION_1_0,
		ULONG(numFonts),
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, header)
	bytes := buf.Bytes()
	if _, err := w.WriteAt(bytes, 0); err != nil {
		return 0, 0, err
	}
	return calcCheckSum(bytes), int64(len(bytes)), nil
}

func writeFontsOffset(w io.WriterAt, offset int64, fonts TTCOffsetTable) (ULONG, error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, fonts)
	bytes := buf.Bytes()
	if _, err := w.WriteAt(bytes, offset); err != nil {
		return 0, err
	}
	return calcCheckSum(bytes), nil
}

func (o OTF) GenerateTTC(w io.WriterAt) error {
	total, offset, err := writeTTCHeader(o, w)
	if err != nil {
		return err
	}
	offsetOffset := offset
	numFonts := len(o)
	fontsOffset := make(TTCOffsetTable, numFonts)
	offset += int64(binary.Size(fontsOffset))
	dictOffset := make([]int64, numFonts)
	entry := make([][]OffsetEntry, numFonts)
	table := make([]Table, 0)
	tableSet := make(map[Table]bool)
	for i, f := range o {
		fontsOffset[i] = ULONG(offset)
		t, d, err := writeSfntHeader(f, w, offset)
		if err != nil {
			return err
		}
		total += t
		dictOffset[i] = d
		numTables := len(f.Table)
		e := make([]OffsetEntry, numTables)
		entry[i] = e
		offset = roundUp(d + int64(binary.Size(e)))
		for _, v := range f.Table {
			if _, ok := tableSet[v]; !ok {
				tableSet[v] = true
				table = append(table, v)
			}
		}
	}
	if t, err := writeFontsOffset(w, offsetOffset, fontsOffset); err == nil {
		total += t
	} else {
		return err
	}
	entryMap, head, total, err := writeTables(table, w, offset)
	if err != nil {
		return err
	}
	// write table directory
	for i, f := range o {
		tag := make(sort.StringSlice, 0)
		for _, v := range f.Table {
			t := v.Tag()
			tag = append(tag, string(t[:]))
		}
		sort.Sort(tag)
		for j, ts := range tag {
			entry[i][j] = *(entryMap[f.Table[ts]])
		}
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, entry[i])
		bytes := buf.Bytes()
		if _, err := w.WriteAt(bytes, dictOffset[i]); err != nil {
			return err
		}
		total += calcCheckSum(bytes)
	}
	checkSumAdjust(w, head, total)
	return nil
}
