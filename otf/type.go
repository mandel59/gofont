package otf

// Open Font Format Standard section 4.3
type (
	BYTE         uint8
	CHAR         int8
	USHORT       uint16
	SHORT        int16
	UINT24       [3]uint8
	ULONG        uint32
	LONG         int32
	FIXED        int32
	FWORD        int16
	UFWORD       uint16
	F2DOT14      int16
	LONGDATETIME int64
	TAG          [4]uint8
	GlyphID      uint16
	Offset       uint16
)

type (
	PlatformID USHORT
	EncodingID USHORT
	LanguageID USHORT
	NameID     USHORT
)
