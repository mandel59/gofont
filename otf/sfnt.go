package otf

type SFNT struct {
	Header SfntHeader
	Table  map[string]Table
}

const checkSumAdjustmentMagic = 0xB1B0AFBA
