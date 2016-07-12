package ihex

import (
	"encoding/hex"
	"github.com/tejainece/hexutils"
	//"log"
)

type (
	Record struct {
		bytes []byte

		length uint8

		addr uint16

		recType RecType
	}
)

//CalcChksum calculates checksum for the given byte array
func CalcChksum(aBytes []byte) uint8 {
	var lRet uint8

	for _, cData := range aBytes {
		lRet += uint8(cData)
	}

	lRet = -lRet

	return lRet
}

//MakeRecord creates record from the record's content
func MakeRecord(aIn []byte) (*Record, error) {
	lInLen := len(aIn)

	//If length is too short, bail
	if lInLen < RecMinLen {
		return nil, ErrInvalidRecLen
	}

	//Check start character
	if aIn[0] != byte(':') {
		return nil, ErrInvalidRecStartCode
	}

	var lErr error

	//Get length
	lHexLen := uint8(0)
	{
		bTempBuf := [2]byte{}
		copy(bTempBuf[:], aIn[1:3])
		lHexLen, lErr = hexutils.ToUInt8(bTempBuf)
		if lErr != nil {
			return nil, ErrNonHexContent
		}
	}

	//Test length
	if lInLen != int(RecMinLen+(lHexLen*2)) {
		return nil, ErrInvalidRecLen
	}

	//Get record type
	var lType RecType
	{
		var bTypeInt uint8
		bTempBuf := [2]byte{}
		copy(bTempBuf[:], aIn[7:9])
		bTypeInt, lErr = hexutils.ToUInt8(bTempBuf)
		if lErr != nil {
			return nil, ErrNonHexContent
		}
		lType = RecType(int(bTypeInt))
	}

	//Test type
	if lType > LastValidRecType {
		return nil, ErrInvalidRecType
	}

	//Check checksum
	{
		var lStChksum uint8
		bTempBuf := [2]byte{}
		copy(bTempBuf[:], aIn[len(aIn)-2:])
		lStChksum, lErr = hexutils.ToUInt8(bTempBuf)
		//log.Printf("Stored checksum: %x", lStChksum)
		if lErr != nil {
			return nil, ErrNonHexContent
		}

		bTempBuf2 := make([]byte, (len(aIn)-3)/2)
		var bNum int
		bNum, lErr = hex.Decode(bTempBuf2, aIn[1:len(aIn)-2])

		if lErr != nil {
			return nil, ErrNonHexContent
		}

		if bNum != len(bTempBuf2) {
			return nil, ErrNonHexContent
		}

		/*
			for _, cData := range bTempBuf2 {
				log.Printf("%x ", cData)
			}
		*/

		if CalcChksum(bTempBuf2) != lStChksum {
			return nil, ErrInvalidChksum
		}
	}

	var lAddr uint16
	{
		bTempBuf := [4]byte{}
		copy(bTempBuf[:], aIn[3:7])
		lAddr, lErr = hexutils.ToUInt16(bTempBuf)
		if lErr != nil {
			return nil, lErr
		}
	}

	lBytes := make([]byte, len(aIn))
	copy(lBytes, aIn)

	return &Record{
		bytes:   lBytes,
		length:  lHexLen,
		addr:    lAddr,
		recType: lType,
	}, nil
}

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
