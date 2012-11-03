package otf

import (
	"io"
)

func hheaParser(_ SFNT, r *io.SectionReader) Table {
	return ReadTable(r, new(Hhea))
}

func (_ *Hhea) Tag() TAG {
	return TAG_HHEA
}

func (t *Hhea) Bytes() []byte {
	return DumpBigEndian(t)
}

func (t *Hhea) SetUp(f SFNT) error {
	if hmtx, ok := f["hmtx"].(*Hmtx); ok {
		t.NumberOfHMetrics = hmtx.NumberOfHMetrics
	}
	// FIXME: it must be implemented
	return nil
}

type Hhea struct {
	Version             FIXED
	Ascender            FWORD
	Descender           FWORD
	LineGap             FWORD
	AdvanceWidthMax     UFWORD
	MinLeftSideBearing  FWORD
	MinRightSideBearing FWORD
	XMaxExtent          FWORD
	CaretSlopeRise      SHORT
	CaretSlopeRun       SHORT
	CaretOffset         SHORT
	RESERVED            [4]SHORT
	MetricDataFormat    SHORT
	NumberOfHMetrics    USHORT
}
