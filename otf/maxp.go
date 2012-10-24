package otf

type Maxp_0_5 struct {
	Version   FIXED
	NumGlyphs USHORT
}

type Maxp_1_0 struct {
	Maxp_0_5
	MaxPoints             USHORT
	MaxContours           USHORT
	MaxCompositePoints    USHORT
	MaxCompositeContours  USHORT
	MaxZones              USHORT
	MaxTwilightPoints     USHORT
	MaxStorage            USHORT
	MaxFunctionDefs       USHORT
	MaxInstructionDefs    USHORT
	MaxStackElements      USHORT
	MaxSizeOfInstructions USHORT
	MaxComponentElements  USHORT
	MaxComponentDepth     USHORT
}
