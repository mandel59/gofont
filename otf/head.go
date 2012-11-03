package otf

import (
	"io"
)

func headParser(_ SFNT, r *io.SectionReader) Table {
	return ReadTable(r, new(Head))
}

func (_ *Head) Tag() TAG {
	return TAG_HEAD
}

func (t *Head) Bytes() []byte {
	return DumpBigEndian(t)
}

func (t *Head) SetUp(f SFNT) error {
	t.CheckSumAdjustment = 0
	// FIXME: it must be implemented
	return nil
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

const HEAD_MAGIC_NUMBER ULONG = 0x5F0F3CF5
