package otf

import "bytes"

type TableRecord struct {
	Tag      TAG
	CheckSum ULONG
	Offset   ULONG
	Length   ULONG
}

type sliceTableRecord []*TableRecord
type byTagSort struct{ sliceTableRecord }
type byOffsetSort struct{ sliceTableRecord }

func (s sliceTableRecord) Len() int      { return len(s) }
func (s sliceTableRecord) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s byTagSort) Less(i, j int) bool {
	return bytes.Compare(s.sliceTableRecord[i].Tag[:], s.sliceTableRecord[j].Tag[:]) < 0
}

func (s byOffsetSort) Less(i, j int) bool {
	return s.sliceTableRecord[i].Offset < s.sliceTableRecord[j].Offset
}
