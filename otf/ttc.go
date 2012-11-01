package otf

type TTCHeader struct {
	TTCTag   TAG
	Version  FIXED
	NumFonts ULONG
}

type TTCOffsetTable []ULONG

type TTCHeaderDsig struct {
	UIDsigTag    ULONG
	UIDsigLength ULONG
	UIDsigOffset ULONG
}

var TAG_TTC TAG = TAG{'t', 't', 'c', 'f'}
