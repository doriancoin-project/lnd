package chainreg

// ChainCode is an enum-like structure for keeping track of the chains
// currently supported within lnd.
type ChainCode uint32

const (
	// DoriancoinChain is Doriancoin's chain.
	DoriancoinChain ChainCode = 1

	// LitecoinChain is an alias for backwards compatibility.
	LitecoinChain = DoriancoinChain
)

// String returns a string representation of the target ChainCode.
func (c ChainCode) String() string {
	switch c {
	case DoriancoinChain:
		return "doriancoin"
	default:
		return "kekcoin"
	}
}
