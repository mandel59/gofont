package otf

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
)

func parseTTC(blob []byte) (OTF, error) {
	buf := bytes.NewBuffer(blob)
	ttcHeader := new(TTCHeader)
	err := binary.Read(buf, binary.BigEndian, ttcHeader)
	if err != nil {
		return nil, err
	}
	numFonts := ttcHeader.NumFonts
	offsetTable := make(TTCOffsetTable, numFonts)
	if err := binary.Read(buf, binary.BigEndian, offsetTable); err != nil {
		return nil, err
	}
	o := make([]*SFNT, numFonts)
	for i, offset := range offsetTable {
		o[i], err = parseSFNT(blob[offset:], blob)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

func parseSFNT(header, blob []byte) (*SFNT, error) {
	o := new(SFNT)
	buf := bytes.NewBuffer(header)
	err := binary.Read(buf, binary.BigEndian, &o.SfntHeader)
	if err != nil {
		return nil, err
	}
	numTables := o.SfntHeader.NumTables
	offsetTable := make([]OffsetEntry, numTables)
	if err := binary.Read(buf, binary.BigEndian, offsetTable); err != nil {
		return nil, err
	}
	table := make(map[string]Table)
	for _, entry := range offsetTable {
		start := entry.Offset
		end := start + entry.Length
		table[string(entry.Tag[:])] = NewTable(entry.Tag, blob[start:end])
	}
	o.Tables = table
	return o, nil
}

func ReadOTF(r io.Reader) (OTF, error) {
	blob, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if string(blob[0:4]) == string(TAG_TTC[:]) {
		return parseTTC(blob)
	}
	o, err := parseSFNT(blob, blob)
	if err != nil {
		return nil, err
	}
	return []*SFNT{o}, nil
}
