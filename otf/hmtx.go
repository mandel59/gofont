package otf

type LongHorMetric struct {
	AdvanceWidth USHORT
	Lsb          SHORT
}

type Hmtx struct {
	HMetrics        []LongHorMetric
	LeftSideBearing []SHORT
}
