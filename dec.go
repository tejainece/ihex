package ihex

import (
	"encoding/hex"
)

func (meRec Record) GetData() (uint16, []byte, error) {
	//Check type
	if meRec.recType != RecTypeData {
		return 0, nil, ErrInvalidRecType
	}

	lRet := make([]byte, meRec.length)

	lConved, lErr := hex.Decode(lRet, meRec.bytes[9:len(meRec.bytes)-2])

	if lConved != len(lRet) || lErr != nil {
		return 0, nil, ErrNonHexContent
	}

	return meRec.addr, lRet, nil
}
