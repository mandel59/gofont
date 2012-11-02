package otf

import "io"

// OTF attaches the methods of Interface to []*SFNT
type OTF []*SFNT

func (o OTF) Generate(w io.WriterAt) error {
	if len(o) == 1 {
		return o[0].Generate(w)
	}
	header := TTCHeader{
		TAG_TTC,
		VERSION_1_0,
		ULONG(len(o)),
	}
	return o.GenerateTTC(w, header)
}
