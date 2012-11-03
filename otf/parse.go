package otf

import (
	"encoding/binary"
	"io"
)

type TableParserFunc func(f SFNT, r *io.SectionReader) Table
type TableParser map[TAG]TableParserFunc

var stage0 = TableParser{
	TAG_CMAP: cmapParser,
	TAG_HEAD: headParser,
	TAG_OS_2: os_2Parser,
	TAG_HHEA: hheaParser,
}

var stage1 = TableParser{
	//TAG_HMTX: hmtxParser,
}

var DefaultParser = []TableParser {
	stage0,
	stage1,
}

func (parser TableParser) Parse(f SFNT, t Table) Table {
	tag := t.Tag()
	if parser, ok := parser[tag]; ok {
		if d, ok := t.(*DefaultTable); ok {
			r := d.Reader()
			r.Seek(0, 0)
			if p := parser(f, r); p != nil {
				return p
			}
		}
	}
	return t
}

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
	o := make([]SFNT, numFonts)
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

func parseSFNT(r io.ReaderAt, headerOffset int64, table map[int64]Table) (SFNT, error) {
	header := new(SfntHeader)
	headerSize := int64(binary.Size(header))
	sr := io.NewSectionReader(r, headerOffset, headerSize)
	if err := binary.Read(sr, binary.BigEndian, header); err != nil {
		return nil, err
	}
	numTables := header.NumTables
	offsetTable := make([]OffsetEntry, numTables)
	sr = io.NewSectionReader(r, headerOffset+headerSize, int64(binary.Size(offsetTable)))
	if err := binary.Read(sr, binary.BigEndian, offsetTable); err != nil {
		return nil, err
	}
	tableMap := make(SFNT)
	for _, entry := range offsetTable {
		tag := entry.Tag.String()
		offset := int64(entry.Offset)
		size := int64(entry.Length)
		if v, ok := table[offset]; ok {
			tableMap[tag] = v
		} else {
			v = &DefaultTable{entry.Tag, io.NewSectionReader(r, offset, size)}
			table[offset] = v
			tableMap[tag] = v
		}
	}
	for _, p := range DefaultParser {
		for i, v := range tableMap {
			tableMap[i] = p.Parse(tableMap, v)
		}
	}
	return tableMap, nil
}

// Read reads Open Font File from io.ReaderAt.
// It returns OTF and any read error encountered.
func Read(r io.ReaderAt) (OTF, error) {
	tag := make([]byte, 4)
	if _, err := r.ReadAt(tag, 0); err != nil {
		return nil, err
	}
	if string(tag) == TAG_TTC.String() {
		return parseTTC(r)
	}
	o, err := parseSFNT(r, 0, make(map[int64]Table))
	if err != nil {
		return nil, err
	}
	return OTF{o}, nil
}
