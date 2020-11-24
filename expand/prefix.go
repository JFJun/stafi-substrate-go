package expand

import (
	"github.com/JFJun/go-substrate-crypto/ss58"
	"strings"
)

func GetPrefix(chainName string) []byte {
	switch strings.ToLower(chainName) {
	case "polkadot":
		return ss58.PolkadotPrefix
	case "chainx":
		return ss58.ChainXPrefix
	case "kusama":
		return ss58.KsmPrefix
	case "crab":
		return ss58.SubstratePrefix
	case "darwinia":
		return ss58.DarwiniaPrefix
	default:
		return ss58.SubstratePrefix
	}
}
