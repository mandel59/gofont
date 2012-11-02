package otf

import (
	"bytes"
	"encoding/binary"
	"io"
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

func writeTTCHeader(o OTF, w io.WriterAt, header TTCHeader) (ULONG, int64, error) {
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

var binarySizeOffsetEntry = binary.Size(OffsetEntry{})

func (o OTF) GenerateTTC(w io.WriterAt, header TTCHeader) error {
	total, offset, err := writeTTCHeader(o, w, header)
	if err != nil {
		return err
	}
	offsetOffset := offset
	numFonts := len(o)
	fontsOffset := make(TTCOffsetTable, numFonts)
	offset += int64(binary.Size(fontsOffset))
	dictOffset := make([]int64, numFonts)
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
		offset = roundUp(d + int64(len(f.Table) * binarySizeOffsetEntry))
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
	for i, f := range o {
		cs, err := writeTableDirectory(f, entryMap, w, dictOffset[i])
		if err != nil {
			return err
		}
		total += cs
	}
	checkSumAdjust(w, head, total)
	return nil
}
