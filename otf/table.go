package otf

type Table interface {
	Bytes() []byte
	Tag() TAG
	SetUp(f SFNT) error
}

type Subtable interface {
	Bytes() []byte
	Size() int64
}
