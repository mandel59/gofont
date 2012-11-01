package otf

type SFNT struct {
	SfntHeader SfntHeader
	Tables     map[string]Table
}

const checkSumAdjustmentMagic = 0xB1B0AFBA
