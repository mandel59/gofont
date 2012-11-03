package otf

type Cmap struct {
	Version  USHORT
	Subtable []CmapSubtable
}

type CmapHeader struct {
	Version   USHORT
	NumTables USHORT
}

type EncodingRecord struct {
	PlatformID
	EncodingID
	Offset ULONG
}

type CmapSubtable struct {
	PlatformID
	EncodingID
	Subtable
}

var TAG_CMAP = TAG{'c', 'm', 'a', 'p'}
