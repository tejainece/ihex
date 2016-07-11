package ihex

import (
	"github.com/tejainece/hexutils"
)

type (
	Decoder struct {
		baseAddr uint16
	}
)

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (meDec *Decoder) DecodeRecord(aRec Record) (uint32, []byte, error) {
	var lErr error
	switch aRec.recType {
	case RecTypeData:
		{
			var bOffset uint16
			var bBytes []byte

			bOffset, bBytes, lErr = aRec.GetData()

			if lErr != nil {
				return 0, nil, lErr
			}

			bAddr := hexutils.MakeUInt32FromUInt16(meDec.baseAddr, bOffset)

			return bAddr, bBytes, nil
		}
		break
	case RecTypeExtLinAddr:
		{
			if aRec.length != ExtLinAddrLen {
				return 0, nil, ErrInvalidRecLen
			}

			var bAddr uint16
			bTempBuf := [4]byte{}
			copy(bTempBuf[:], aRec.bytes[9:13])
			bAddr, lErr = hexutils.ToUInt16(bTempBuf)
			if lErr != nil {
				return 0, nil, ErrNonHexContent
			}
			meDec.baseAddr = bAddr

			return 0, []byte{}, nil
		}
		break
	case RecTypeExtSegAddr:
	case RecTypeStartSegAddr:
		return 0, nil, ErrSegAddressingNotSupported
		break
	case RecTypeEOF:
		return 0, nil, nil
		break
	}

	return 0, []byte{}, nil
}
