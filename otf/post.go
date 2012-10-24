package otf

type Post struct {
	Version            FIXED
	ItalicAngle        FIXED
	UnderlinePosition  FWORD
	UnderlineThickness FWORD
	IsFixedPitch       ULONG
	MinMemType42       ULONG
	MaxMemType42       ULONG
	MinMemType1        ULONG
	MaxMemType1        ULONG
}

type Post_2_0 struct {
	Post
	NumberOfGlyphs USHORT
	GlyphNameIndex []USHORT
	Names          []CHAR
}
