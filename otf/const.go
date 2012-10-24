package otf

const (
	VERSION_0_5 FIXED = 0x00005000
	VERSION_1_0 FIXED = 0x00010000
	VERSION_2_0 FIXED = 0x00020000
)

// Platform IDs
const (
	PLATFORM_UNICODE = iota
	PLATFORM_MACINTOSH
	PLATFORM_ISO
	PLATFORM_WINDOWS
	PLATFORM_CUSTOM
)

// Unicode platform-specific encoding IDs (platform ID = 0)
const (
	UNICODE_1_0 = iota
	UNICODE_1_1
	ISO_IEC_10646
	UNICODE_BMP
	UNICODE_FULL
	UNICODE_VS
)

// Windows platform-specific encoding IDs (platform ID = 3)
const (
	WINDOWS_SYMBOL = iota
	WINDOWS_UCS2
	WINDOWS_SJIS
	WINDOWS_PRC
	WINDOWS_BIG5
	WINDOWS_WANSUNG
	WINDOWS_JOHAB
	_
	_
	_
	WINDOWS_UCS4
)

// Macintosh platform-specific encoding IDs (platform ID = 1)
const (
	MACINTOSH_ROMAN = iota
	MACINTOSH_JAPANESE
	MACINTOSH_CHINESE_TRAD
	MACINTOSH_KOREAN
	MACINTOSH_ARBIC
	MACINTOSH_HEBREW
	MACINTOSH_GREEK
	MACINTOSH_RUSSIAN
	MACINTOSH_RSYMBOL
	MACINTOSH_DEVANAGARI
	MACINTOSH_GURMUKHI
	MACINTOSH_GUJARATI
	MACINTOSH_ORIYA
	MACINTOSH_BENGALI
	MACINTOSH_TAMIL
	MACINTOSH_TELUGU
	MACINTOSH_KANNADA
	MACINTOSH_MALAYALAM
	MACINTOSH_SINHALESE
	MACINTOSH_BURMESE
	MACINTOSH_KHMER
	MACINTOSH_THAI
	MACINTOSH_LAOTIAN
	MACINTOSH_GEORGIAN
	MACINTOSH_ARMENIAN
	MACINTOSH_CHINESE_SIMP
	MACINTOSH_TIBETAN
	MACINTOSH_MONGOLIAN
	MACINTOSH_SLAVIC
	MACINTOSH_VIETNAMESE
	MACINTOSH_SINDHI
	MACINTOSH_UNINTERPRETED
)
