package otf

import "encoding/binary"

func roundUp(n int) int {
	return (n + 3) &^ 3
}

func pad(b []byte) []byte {
	n := len(b)
	l := roundUp(n) - n
	return make([]byte, l)
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
