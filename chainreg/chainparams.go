package chainreg

import (
	"github.com/ltcsuite/lnd/keychain"
	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcd/wire"
)

// DoriancoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type DoriancoinNetParams struct {
	*chaincfg.Params
	RPCPort  string
	CoinType uint32
}

// LitecoinNetParams is an alias for backwards compatibility.
type LitecoinNetParams = DoriancoinNetParams

// DoriancoinSimNetParams contains parameters specific to the simulation test
// network.
var DoriancoinSimNetParams = DoriancoinNetParams{
	Params:   &chaincfg.TestNet4Params,
	RPCPort:  "19556",
	CoinType: keychain.CoinTypeTestnet,
}

// LitecoinSimNetParams is an alias for backwards compatibility.
var LitecoinSimNetParams = DoriancoinSimNetParams

// DoriancoinTestNetParams contains parameters specific to the 4th version of the
// test network.
var DoriancoinTestNetParams = DoriancoinNetParams{
	Params:   &chaincfg.TestNet4Params,
	RPCPort:  "19334",
	CoinType: keychain.CoinTypeTestnet,
}

// LitecoinTestNetParams is an alias for backwards compatibility.
var LitecoinTestNetParams = DoriancoinTestNetParams

// DoriancoinMainNetParams contains the parameters specific to the current
// Doriancoin mainnet.
var DoriancoinMainNetParams = DoriancoinNetParams{
	Params:   &chaincfg.MainNetParams,
	RPCPort:  "1948",
	CoinType: keychain.CoinTypeDoriancoin,
}

// LitecoinMainNetParams is an alias for backwards compatibility.
var LitecoinMainNetParams = DoriancoinMainNetParams

// DoriancoinRegTestNetParams contains parameters specific to a local doriancoin
// regtest network.
var DoriancoinRegTestNetParams = DoriancoinNetParams{
	Params:   &chaincfg.RegressionNetParams,
	RPCPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// LitecoinRegTestNetParams is an alias for backwards compatibility.
var LitecoinRegTestNetParams = DoriancoinRegTestNetParams

// IsTestnet tests if the givern params correspond to a testnet
// parameter configuration.
func IsTestnet(params *DoriancoinNetParams) bool {
	switch params.Params.Net {
	case wire.TestNet4:
		return true
	default:
		return false
	}
}
