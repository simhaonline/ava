package validator

// Qtum ...
type Qtum struct{}

var _ validatorBitcoinLike = (*Qtum)(nil)

// ValidateAddress returns validate result of qtum address
func (v *Qtum) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := normalAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}

	return &Result{IsValid: false, Type: Unknown}
}

// AddressVersion returns qtum address version according to the address type and
// network type
func (v *Qtum) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return 58
		}
		return 50
	}

	if addrType == P2PKH {
		return 120
	}
	return 110
}

// AddressHrp returns hrps of qtum according to the network
func (v *Qtum) AddressHrp(network NetworkType) string {
	panic(ErrUnsupported.Error())
}

// SegwitProgramLength returns segwit program length of qtum
func (v *Qtum) SegwitProgramLength(addrType AddressType) int {
	panic(ErrUnsupported.Error())
}
