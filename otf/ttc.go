package otf

type TTCHeader_1_0 struct {
	TTCTag      TAG
	Version     FIXED
	NumFonts    ULONG
	OffsetTable []ULONG
}

type TTCHeader_2_0 struct {
	TTCHeader_1_0
	UIDsigTag    ULONG
	UIDsigLength ULONG
	UIDsigOffset ULONG
}

var TAG_TTC TAG = TAG{'t', 't', 'c', 'f'}
