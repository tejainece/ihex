package ihex

import (
	"gopkg.in/check.v1"
	"testing"
)

// Test hooks up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

//TstRecord is the test suit
type TstRecord struct{}

var _ = check.Suite(&TstRecord{})

//TestChecksum tests checksum calculation
func (s *TstRecord) TestChecksum(c *check.C) {
	lBytes := []byte{0x03, 0x00, 0x30, 0x00, 0x02, 0x33, 0x7A}
	c.Check(CalcChksum(lBytes), check.Equals, uint8(0x1E))
}

//TestChecksum tests
func (s *TstRecord) TestMakeRecord_1(c *check.C) {
	lBytes := []byte(":0300300002337A1E")

	_, lErr := MakeRecord(lBytes)

	c.Check(lErr, check.IsNil)
	//TODO
}

//TestChecksum tests
func (s *TstRecord) TestMakeRecord_2(c *check.C) {
	lBytes := []byte(":10010000214601360121470136007EFE09D2190140")

	_, lErr := MakeRecord(lBytes)

	c.Check(lErr, check.IsNil)
	//TODO
}

/*
//TestFindLevel tests hello world
func (s *MySuite) TestGetDataFromRecordInvalidStartCode(c *check.C) {
	//_, lErr := GetDataFromRecord([]byte(":050100000"))
	//c.Check(lErr, check.Equals, ErrInvalidRecLen)
}

//TestFindLevel tests hello world
func (s *MySuite) TestGetDataFromRecordInvalidStartCode(c *check.C) {
	_, lErr := GetDataFromRecord([]byte("<05010000065756"))
	c.Check(lErr, check.Equals, ErrInvalidRecStartCode)
}
*/
