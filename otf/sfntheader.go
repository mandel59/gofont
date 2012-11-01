package otf

func (t *SfntHeader) SetNumTables(numTables USHORT) {
	e := USHORT(maxPowerOf2(numTables))
	n := USHORT(1 << e)
	t.NumTables = numTables
	t.SearchRange = n * 16
	t.EntrySelector = e
	t.RangeShift = (numTables - n) * 16
}

type SfntHeader struct {
	SfntVersion   FIXED
	NumTables     USHORT
	SearchRange   USHORT
	EntrySelector USHORT
	RangeShift    USHORT
}

const VERSION_CFF FIXED = 0x4F54544F // "OTTO"
