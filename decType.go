package ihex

type RecType int

const (
	RecTypeData         = 0x00
	RecTypeEOF          = 0x01
	RecTypeExtSegAddr   = 0x02
	RecTypeStartSegAddr = 0x03
	RecTypeExtLinAddr   = 0x04
	RecTypeStartLinAddr = 0x05

	LastValidRecType = RecTypeStartLinAddr
)

const (
	RecMinLen     = 11
	ExtLinAddrLen = 2
)
