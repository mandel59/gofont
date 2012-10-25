package otf

import "errors"

func (t *OffsetTable) VerifyOffsetTable() error {
	n := t.NumTables.MaxPowerOf2()
	entrySelector, _ := n.Log2()
	if t.SearchRange == n*16 && t.EntrySelector == entrySelector && t.RangeShift == (t.NumTables-n)*16 {
		return nil
	}
	return errors.New("VerifyOffsetTable: Invalid offset table")
}

func (t *OffsetTable) SetupNumTables() {
	numTables := t.NumTables
	n := numTables.MaxPowerOf2()
	t.SearchRange = n * 16
	t.EntrySelector, _ = n.Log2()
	t.RangeShift = (numTables - n) * 16
}

type OffsetTable struct {
	SfntVersion   FIXED
	NumTables     USHORT
	SearchRange   USHORT
	EntrySelector USHORT
	RangeShift    USHORT
}

const VERSION_CFF FIXED = 0x4F54544F // "OTTO"
