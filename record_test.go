package ihex

import (
	"gopkg.in/check.v1"
	"testing"
)

// Test hooks up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

//MySuite is the test suit
type MySuite struct{}

var _ = check.Suite(&MySuite{})

/*
//TestFindLevel tests hello world
func (s *MySuite) TestGetDataFromRecordInvalidLen(c *check.C) {
	_, lErr := GetDataFromRecord([]byte(":050100000"))
	c.Check(lErr, check.Equals, ErrInvalidRecLen)
}

//TestFindLevel tests hello world
func (s *MySuite) TestGetDataFromRecordInvalidStartCode(c *check.C) {
	_, lErr := GetDataFromRecord([]byte("<05010000065756"))
	c.Check(lErr, check.Equals, ErrInvalidRecStartCode)
}
*/
