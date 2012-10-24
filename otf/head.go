package otf

import (
	"encoding/binary"
	"io"
)

func (_ *Head) Tag() TAG {
	return TAG_HEAD
}

func (_ *Head) CheckSum() ULONG {
	// FIXME: not impremented
	return 0;
}

func (t *Head) Len() ULONG {
	return ULONG(binary.Size(t))
}

func (t *Head) WriteTo(w io.Writer) (n int, err error) {
	n = binary.Size(t)
	err = binary.Write(w, binary.BigEndian, t)
	return
}

type Head struct {
	Version            FIXED
	FontRevision       FIXED
	CheckSumAdjustment ULONG
	MagicNumber        ULONG
	Flags              USHORT
	UnitsPerEm         USHORT
	Created            LONGDATETIME
	Modified           LONGDATETIME
	XMin               SHORT
	YMin               SHORT
	XMax               SHORT
	YMax               SHORT
	MacStyle           USHORT
	LowestRecPPEM      USHORT
	FontDirectionHint  SHORT
	IndexToLocFormat   SHORT
	GlyphDataFormat    SHORT
}

var TAG_HEAD = TAG{'h', 'e', 'a', 'd'}
const HEAD_MAGIC_NUMBER ULONG = 0x5F0F3CF5
