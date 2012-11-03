package otf

import (
	"io"
)

func hmtxParser(f SFNT, r *io.SectionReader) Table {
	if t, ok := f["hhea"]; ok {
		if hhea, ok := t.(*Hhea); ok {
			hmtx := &Hmtx{
				&SubtableReader{r},
				hhea.NumberOfHMetrics,
			}
			return hmtx
		}
	}
	return nil
}

func (_ *Hmtx) Tag() TAG {
	return TAG_HMTX
}

func (t *Hmtx) SetUp(f SFNT) error {
	return nil
}

type LongHorMetric struct {
	AdvanceWidth    UFWORD
	LeftSideBearing FWORD
}

type Hmtx struct {
	Subtable
	NumberOfHMetrics USHORT
}
