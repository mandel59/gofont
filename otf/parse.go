package otf

import (
	"encoding/binary"
	"sort"
	"io"
)

func ReadOTF(r io.Reader) (otf *OTF, err error) {
	otf = new(OTF)
	err = binary.Read(r, binary.BigEndian, &otf.OffsetTable)
	if err != nil {
		return
	}
	err = otf.OffsetTable.VerifyOffsetTable()
	if err != nil {
		return
	}
	otf.TableRecords = make(sliceTableRecord, otf.NumTables)
	numTables := int(otf.NumTables)
	for i := 0; i < numTables; i++ {
		rec := new(TableRecord)
		binary.Read(r, binary.BigEndian, rec)
		otf.TableRecords[i] = rec
	}
	m := ULONG(0)
	for _, r := range otf.TableRecords {
		n := r.Offset + r.Length
		if m < n {
			m = n
		}
	}
	offset := binary.Size(otf.OffsetTable) + binary.Size(otf.TableRecords);
	sort.Sort(byOffsetSort{otf.TableRecords})
	blob := make([]byte, int(m) - offset)
	_, err = r.Read(blob)
	if err != nil {
		return
	}
	otf.Tables = make([]Table, otf.NumTables)
	for i, r := range otf.TableRecords {
		start := ULONG(int(r.Offset) - offset)
		end := start + r.Length
		otf.Tables[i] = NewTable(r.Tag, blob[start:end])
	}
	return
}
