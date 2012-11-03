package otf

import (
	"io"
	"io/ioutil"
)

type DefaultTable struct {
	tag    TAG
	reader *io.SectionReader
}

func (t *DefaultTable) Tag() TAG {
	return t.tag
}

func (t *DefaultTable) Bytes() []byte {
	t.reader.Seek(0, 0)
	bytes, _ := ioutil.ReadAll(t.reader)
	return bytes
}

func (_ *DefaultTable) SetUp(f SFNT) error {
	return nil
}
