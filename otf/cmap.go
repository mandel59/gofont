package otf

type Cmap struct {
	CmapHeader
	EncodingRecord []EncodingRecord
	Subtables      []CmapSubtable
}

type CmapSubtable Subtable

type CmapHeader struct {
	Version   USHORT
	NumTables USHORT
}

type EncodingRecord struct {
	PlatformID
	EncodingID
	Offset ULONG
}

type CmapFormat0 struct {
	Format       USHORT
	Length       USHORT
	Language     USHORT
	GlyphIdArray [256]BYTE
}

type CmapFormat2 struct {
	Format          USHORT
	Length          USHORT
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
	Format        USHORT
	Length        USHORT
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
	Format       USHORT
	Length       USHORT
	Language     USHORT
	FirstCode    USHORT
	EntryCount   USHORT
	GlyphIdArray []USHORT
}

type CmapFormat8 struct {
	Format   USHORT
	_        USHORT
	Length   ULONG
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
	Format        USHORT
	_             USHORT
	Length        ULONG
	Language      ULONG
	StartCharCode ULONG
	NumChars      ULONG
	Glyphs        []USHORT
}

type CmapFormat12 struct {
	Format   USHORT
	_        USHORT
	Length   ULONG
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
	Format   USHORT
	_        USHORT
	Length   ULONG
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
