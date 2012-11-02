package otf

import (
	"encoding/binary"
	"time"
)

func s_tag(s string) (t TAG) {
	copy(t[0:4], []byte(s[0:4]))
	return
}

func roundUp(n int64) int64 {
	return (n + 3) &^ 3
}

func maxPowerOf2(n USHORT) USHORT {
	e := 0
	for n >>= 1; n > 0; n >>= 1 {
		e++
	}
	return USHORT(e)
}

func calcCheckSum(t []byte) (s ULONG) {
	for i := 0; i < len(t); i += 4 {
		if i+4 > len(t) {
			var u [4]byte
			copy(u[0:len(t)-i], t[i:len(t)])
			s += ULONG(binary.BigEndian.Uint32(u[:]))
		} else {
			s += ULONG(binary.BigEndian.Uint32(t[i : i+4]))
		}
	}
	return
}

var epoch = time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC)

func longDateTime(t time.Time) LONGDATETIME {
	return LONGDATETIME(t.Sub(epoch) / time.Second)
}
