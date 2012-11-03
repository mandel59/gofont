package otf

import (
	"io"
)

type TableParserFunc func(r *io.SectionReader) Table
type TableParser map[TAG]TableParserFunc

var DefaultParser = TableParser{
	TAG_CMAP: cmapParser,
	TAG_HEAD: headParser,
	TAG_OS_2: os_2Parser,
}

func (parser TableParser) Parse(t TAG, r *io.SectionReader) Table {
	if parser, ok := parser[t]; ok {
		return parser(r)
	}
	return nil
}

func NewTable(t TAG, r *io.SectionReader) Table {
	table := DefaultParser.Parse(t, r)
	if table != nil {
		return table
	}
	return &DefaultTable{t, r}
}

type Table interface {
	Subtable
	Tag() TAG
	SetUp(f SFNT) error
}

type Subtable interface {
	Bytes() []byte
}
