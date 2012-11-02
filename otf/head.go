package otf

import (
	"bytes"
	"encoding/binary"
)

func (_ *Head) Tag() TAG {
	return TAG_HEAD
}

func (t *Head) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, t)
	return buf.Bytes()
}

func (t *Head) Set() {
	t.CheckSumAdjustment = 0
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
