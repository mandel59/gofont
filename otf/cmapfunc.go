package otf

import (
	"bytes"
	"encoding/binary"
	"io"
)

func cmapParser(_ SFNT, r *io.SectionReader) Table {
	t := new(Cmap)
	header := CmapHeader{}
	if err := binary.Read(r, binary.BigEndian, &header); err != nil {
		return nil
	}
	t.Version = header.Version
	numTables := header.NumTables
	encodingRecord := make([]EncodingRecord, numTables)
	if err := binary.Read(r, binary.BigEndian, encodingRecord); err != nil {
		return nil
	}
	t.Subtable = make([]CmapSubtable, numTables)
	for i, v := range encodingRecord {
		pid, eid, offset := v.PlatformID, v.EncodingID, int64(v.Offset)
		format := make([]byte, 8)
		r.ReadAt(format, offset)
		var length int64
		switch binary.BigEndian.Uint16(format[:2]) {
		case 0, 2, 4, 6:
			length = int64(binary.BigEndian.Uint16(format[2:4]))
			break
		case 8, 10, 12, 13:
			length = int64(binary.BigEndian.Uint32(format[4:8]))
			break
		case 14:
			length = int64(binary.BigEndian.Uint32(format[2:6]))
			break
		default:
			return nil
		}
		sr := io.NewSectionReader(r, offset, length)
		t.Subtable[i] = CmapSubtable{pid, eid, &SubtableReader{sr}}
	}
	return Table(t)
}

func (_ *Cmap) Tag() TAG {
	return TAG_CMAP
}

func (t *Cmap) Bytes() []byte {
	numTables := len(t.Subtable)
	header := CmapHeader{
		t.Version,
		USHORT(numTables),
	}
	encodingRecord := make([]EncodingRecord, numTables)
	p := int64(binary.Size(header) + binary.Size(encodingRecord))
	for i, v := range t.Subtable {
		encodingRecord[i].PlatformID = v.PlatformID
		encodingRecord[i].EncodingID = v.EncodingID
		encodingRecord[i].Offset = ULONG(p)
		p = roundUp(p + int64(v.Size()))
	}
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, header); err != nil {
		return nil
	}
	if err := binary.Write(buf, binary.BigEndian, encodingRecord); err != nil {
		return nil
	}
	for _, v := range t.Subtable {
		bytes := v.Bytes()
		size := int64(len(bytes))
		if err := binary.Write(buf, binary.BigEndian, bytes); err != nil {
			return nil
		}
		if err := binary.Write(buf, binary.BigEndian, make([]byte, roundUp(size)-size)); err != nil {
			return nil
		}
	}
	return buf.Bytes()
}

func (t *Cmap) SetUp(f SFNT) error {
	return nil
}
