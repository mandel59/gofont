package otf

func NewOffsetTable(version FIXED, numTables USHORT) OffsetTable {
	n := numTables.MaxPowerOf2()
	searchRange := n * 16
	entrySelector, _ := n.Log2()
	rangeShift := numTables*16 - searchRange
	return OffsetTable{
		version,
		numTables,
		searchRange,
		entrySelector,
		rangeShift,
	}
}

type OffsetTable struct {
	SfntVersion   FIXED
	NumTables     USHORT
	SearchRange   USHORT
	EntrySelector USHORT
	RangeShift    USHORT
}

const VERSION_CFF FIXED = 0x4F54544F // "OTTO"
