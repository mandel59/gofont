package otf

type OS_2 struct {
	Version             USHORT
	XAvgCharWidth       SHORT
	UsWeightClass       USHORT
	UsWidthClass        USHORT
	FsType              USHORT
	YSubscriptXSize     SHORT
	YSubscriptYSize     SHORT
	YSubscriptXOffset   SHORT
	YSubscriptYOffset   SHORT
	YSuperscriptXSize   SHORT
	YSuperscriptYSize   SHORT
	YSuperscriptXOffset SHORT
	YSuperscriptYOffset SHORT
	YStrikeoutSize      SHORT
	YStrikeoutPosition  SHORT
	SFamilyClass        SHORT
	Panose
	UlUnicodeRange1  ULONG
	UlUnicodeRange2  ULONG
	UlUnicodeRange3  ULONG
	UlUnicodeRange4  ULONG
	AchVendID        [4]CHAR
	FsSelection      USHORT
	UsFirstCharIndex USHORT
	UsLastCharIndex  USHORT
	STypoAscender    SHORT
	STypoDescender   SHORT
	STypoLineGap     SHORT
	UsWinAscent      USHORT
	UsWinDescent     USHORT
	UlCodePageRange1 ULONG
	UlCodePageRange2 ULONG
	SxHeight         SHORT
	SCapHeight       SHORT
	UsDefaultChar    USHORT
	UsBreakChar      USHORT
	UsMaxContext     USHORT
}

type Panose struct {
	BFamilyType      BYTE
	BSerifStyle      BYTE
	BWeight          BYTE
	BProportion      BYTE
	BContrast        BYTE
	BStrokeVariation BYTE
	BArmStyle        BYTE
	BLetterform      BYTE
	BMidline         BYTE
	BXHeight         BYTE
}
