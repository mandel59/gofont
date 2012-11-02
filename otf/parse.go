package otf

import (
	"encoding/binary"
	"io"
)

func parseTTC(r io.ReaderAt) (OTF, error) {
	ttcHeader := new(TTCHeader)
	ttcHeaderSize := int64(binary.Size(ttcHeader))
	sr := io.NewSectionReader(r, 0, ttcHeaderSize)
	if err := binary.Read(sr, binary.BigEndian, ttcHeader); err != nil {
		return nil, err
	}
	numFonts := ttcHeader.NumFonts
	offsetTable := make(TTCOffsetTable, numFonts)
	sr = io.NewSectionReader(r, ttcHeaderSize, int64(binary.Size(offsetTable)))
	if err := binary.Read(sr, binary.BigEndian, offsetTable); err != nil {
		return nil, err
	}
	o := make([]*SFNT, numFonts)
	table := make(map[int64]Table)
	for i, offset := range offsetTable {
		var err error
		o[i], err = parseSFNT(r, int64(offset), table)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

func parseSFNT(r io.ReaderAt, headerOffset int64, table map[int64]Table) (*SFNT, error) {
	o := new(SFNT)
	headerSize := int64(binary.Size(o.Header))
	sr := io.NewSectionReader(r, headerOffset, headerSize)
	if err := binary.Read(sr, binary.BigEndian, &o.Header); err != nil {
		return nil, err
	}
	numTables := o.Header.NumTables
	offsetTable := make([]OffsetEntry, numTables)
	sr = io.NewSectionReader(r, headerOffset+headerSize, int64(binary.Size(offsetTable)))
	if err := binary.Read(sr, binary.BigEndian, offsetTable); err != nil {
		return nil, err
	}
	tableMap := make(map[string]Table)
	for _, entry := range offsetTable {
		tag := string(entry.Tag[:])
		offset := int64(entry.Offset)
		size := int64(entry.Length)
		if v, ok := table[offset]; ok {
			tableMap[tag] = v
		} else {
			v = NewTable(entry.Tag, io.NewSectionReader(r, offset, size))
			table[offset] = v
			tableMap[tag] = v
		}
	}
	o.Table = tableMap
	return o, nil
}

// Read reads Open Font File from io.ReaderAt.
// It returns OTF and any read error encountered.
func Read(r io.ReaderAt) (OTF, error) {
	tag := make([]byte, 4)
	if _, err := r.ReadAt(tag, 0); err != nil {
		return nil, err
	}
	if string(tag) == string(TAG_TTC[:]) {
		return parseTTC(r)
	}
	o, err := parseSFNT(r, 0, make(map[int64]Table))
	if err != nil {
		return nil, err
	}
	return []*SFNT{o}, nil
}
