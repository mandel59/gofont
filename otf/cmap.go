package otf

import "io"

type Cmap struct {
	Version  USHORT
	Subtable []CmapSubtable
}

type CmapSubtable struct {
	PlatformID
	EncodingID
	io.ReadSeeker
}

func (t *CmapSubtable) Size() int64 {
	n, err := t.Seek(0, 2)
	if err != nil {
		return 0
	}
	return n
}

func (t *CmapSubtable) Bytes() []byte {
	bytes := make([]byte, t.Size())
	t.Seek(0, 0)
	t.Read(bytes)
	return bytes
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

type CmapShortSubtable struct {
	Format USHORT
	Length USHORT
}

type CmapLongSubtable struct {
	Format USHORT
	_      USHORT
	Length ULONG
}

type CmapFormat0 struct {
	CmapShortSubtable
	Language     USHORT
	GlyphIdArray [256]BYTE
}

type CmapFormat2 struct {
	CmapShortSubtable
	Language        USHORT
	SubHeaderKeys   [256]USHORT
	SubHeaders      []CmapFormat2SubHeader
	GlyphIndexArray []USHORT
}

type CmapFormat2SubHeader struct {
	FirstCode     USHORT
	EntryCount    USHORT
	IdDelta       SHORT
	IdRangeOffset USHORT
}

type CmapFormat4 struct {
	CmapShortSubtable
	Language      USHORT
	SegCountX2    USHORT
	SearchRange   USHORT
	EntrySelector USHORT
	RangeShift    USHORT
	EndCode       []USHORT
	_             USHORT
	StartCode     []USHORT
	IdDelta       []SHORT
	IdRangeOffset []USHORT
	GlyphIdArray  []USHORT
}

type CmapFormat6 struct {
	CmapShortSubtable
	Language     USHORT
	FirstCode    USHORT
	EntryCount   USHORT
	GlyphIdArray []USHORT
}

type CmapFormat8 struct {
	CmapLongSubtable
	Language ULONG
	Is32     [8192]BYTE
	NGroups  ULONG
	Groups   []CmapFormat8Group
}

type CmapFormat8Group struct {
	StartCharCode ULONG
	EndCharCode   ULONG
	StartGlyphID  ULONG
}

type CmapFormat10 struct {
	CmapLongSubtable
	Language      ULONG
	StartCharCode ULONG
	NumChars      ULONG
	Glyphs        []USHORT
}

type CmapFormat12 struct {
	CmapLongSubtable
	Language ULONG
	NGroups  ULONG
	Groups   []CmapFormat12Group
}

type CmapFormat12Group struct {
	StartCharCode ULONG
	EndCharCode   ULONG
	StartGlyphID  ULONG
}

type CmapFormat13 struct {
	CmapLongSubtable
	Language ULONG
	NGroups  ULONG
	Groups   []CmapFormat13Group
}

type CmapFormat13Group struct {
	StartCharCode ULONG
	EndCharCode   ULONG
	GlyphID       ULONG
}

type CmapFormat14 struct {
	Format   USHORT
	Length   ULONG
	NRecords ULONG
	Records  []VariationSelectorRecord
	DefaultUVSTable
	NonDefaultUVSTable
}

type VariationSelectorRecord struct {
	VarSelector         UINT24
	DefaultUVSOffset    ULONG
	NonDefaultUVSOffset ULONG
}

type DefaultUVSTable struct {
	NRanges ULONG
	Ranges  []UnicodeValueRange
}

type UnicodeValueRange struct {
	StartUnicodeValue UINT24
	AdditionalCount   BYTE
}

type NonDefaultUVSTable struct {
	NMappings ULONG
	Mappings  []UVSMapping
}

type UVSMapping struct {
	UnicodeValue UINT24
	GlyphID      USHORT
}

var TAG_CMAP = TAG{'c', 'm', 'a', 'p'}
