package otf

import "errors"

func (n USHORT) MaxPowerOf2() USHORT {
	l := n
	for i := 0; i < 15; i++ {
		n >>= 1
		l &^= n
	}
	return l
}

func (n USHORT) Log2() (USHORT, error) {
	for i := USHORT(15); i >= 0; i-- {
		if n&(1<<i) != 0 {
			return i, nil
		}
	}
	return 0, errors.New("USHORT Log2: Invarid anti-logarithm")
}

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
