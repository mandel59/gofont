package otf

import "io"

type UnknownTable struct {
	TableTag TAG
	Body     []byte
}

func (t *UnknownTable) Tag() TAG {
	return t.TableTag
}

func (t *UnknownTable) CheckSum() ULONG {
	return calcCheckSum(t.Body)
}

func (t *UnknownTable) Len() ULONG {
	return ULONG(len(t.Body))
}

func (t *UnknownTable) WriteTo(w io.Writer) (n int, err error) {
	return w.Write(t.Body)
}
