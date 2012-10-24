package otf

type NameFormat0 struct {
	Format       USHORT
	Count        USHORT
	StringOffset USHORT
	NameRecord   []NameRecord
}

type NameFormat1 struct {
	NameFormat0
	LangTagCount  USHORT
	LangTagRecord []LangTagRecord
}

type LangTagRecord struct {
	Length USHORT
	Offset USHORT
}

type NameRecord struct {
	PlatformID
	EncodingID
	LanguageID
	NameID
	Length USHORT
	Offset USHORT
}

const (
	COPYRIGHT_NOTICE = iota
	FONT_FAMILY
	FONT_SUBFAMILY
	UNIQUE_FONT_IDENTIFIER
	FULL_FONT_NAME
	VERSION_STRING
	POSTSCRIPT
	TRADEMARK
	MANUFACTURER
	DESIGNER
	DESCRIPTION
	URL_VENDOR
	URL_DESIGNER
	LICENSE_DESCRIPTION
	LICENSE_INFO_URL
	_
	PREFERRED_FAMILY
	PREFERRED_SUBFAMILY
	COMPATIBLE_FULL
	SAMPLE_TEXT
	CID_FINDFONT_NAME
	WWS_FAMILY
	WWS_SUBFAMILY
)
