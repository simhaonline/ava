package validator

// Dash ...
type Dash struct{}

var _ validatorBitcoinLike = (*Dash)(nil)

// ValidateAddress returns validate result of dash address
func (v *Dash) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := normalAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}

	return &Result{IsValid: false, Type: Unknown}
}

// AddressVersion returns dash address version according to the address type and
// network type
func (v *Dash) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return 76
		}
		return 16
	}

	if addrType == P2PKH {
		return 140
	}
	return 19
}

// AddressHrp returns hrps of dash according to the network
func (v *Dash) AddressHrp(network NetworkType) string {
	panic(ErrUnsupported.Error())
}

// SegwitProgramLength returns segwit program length of dash
func (v *Dash) SegwitProgramLength(addrType AddressType) int {
	panic(ErrUnsupported.Error())
}
