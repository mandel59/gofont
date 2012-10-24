package otf

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
	_                   SHORT
	_                   SHORT
	_                   SHORT
	_                   SHORT
	MetricDataFormat    SHORT
	NumberOfHMetrics    USHORT
}
