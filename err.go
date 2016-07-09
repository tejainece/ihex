package ihex

//Err contains information about the failure
type Err string

//Error returns printable string of the failure
func (meErr Err) Error() string {
	return string(meErr)
}

const (
	ErrInvalidRecLen             = Err("Invalid record length!")
	ErrInvalidRecStartCode       = Err("Invalid record start code!")
	ErrNonHexContent             = Err("Encountered non hex content!")
	ErrInvalidRecType            = Err("Invalid record type!")
	ErrInvalidChksum             = Err("Invalid checksum!")
	ErrSegAddressingNotSupported = Err("Segment addressing not supported!")
)
