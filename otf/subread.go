package otf

import "io"

type SubtableReader struct {
	io.ReadSeeker
}

func (t *SubtableReader) Size() int64 {
	n, err := t.Seek(0, 2)
	if err != nil {
		return 0
	}
	return n
}

func (t *SubtableReader) Bytes() []byte {
	bytes := make([]byte, t.Size())
	t.Seek(0, 0)
	t.Read(bytes)
	return bytes
}
