package types

import (
	"bytes"
	"fmt"
	"time"
)

// GenesisState is the state that must be provided at genesis.
type GenesisState struct {
	Params                Params                   `json:"params" yaml:"params"`
	JPYXAccumulationTimes GenesisAccumulationTimes `json:"jpyx_accumulation_times" yaml:"jpyx_accumulation_times"`
	JPYXMintingClaims     JPYXMintingClaims        `json:"jpyx_minting_claims" yaml:"jpyx_minting_claims"`
}

// NewGenesisState returns a new genesis state
func NewGenesisState(params Params, jpyxAccumTimes GenesisAccumulationTimes, c JPYXMintingClaims) GenesisState {
	return GenesisState{
		Params:                params,
		JPYXAccumulationTimes: jpyxAccumTimes,
		JPYXMintingClaims:     c,
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params:                DefaultParams(),
		JPYXAccumulationTimes: GenesisAccumulationTimes{},
		JPYXMintingClaims:     DefaultJPYXClaims,
	}
}

// Validate performs basic validation of genesis data returning an
// error for any failed validation criteria.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	if err := gs.JPYXAccumulationTimes.Validate(); err != nil {
		return err
	}

	return gs.JPYXMintingClaims.Validate()
}

// Equal checks whether two gov GenesisState structs are equivalent
func (gs GenesisState) Equal(gs2 GenesisState) bool {
	b1 := ModuleCdc.MustMarshalBinaryBare(gs)
	b2 := ModuleCdc.MustMarshalBinaryBare(gs2)
	return bytes.Equal(b1, b2)
}

// IsEmpty returns true if a GenesisState is empty
func (gs GenesisState) IsEmpty() bool {
	return gs.Equal(GenesisState{})
}

// GenesisAccumulationTime stores the previous reward distribution time and its corresponding collateral type
type GenesisAccumulationTime struct {
	CollateralType           string    `json:"collateral_type" yaml:"collateral_type"`
	PreviousAccumulationTime time.Time `json:"previous_accumulation_time" yaml:"previous_accumulation_time"`
}

// NewGenesisAccumulationTime returns a new GenesisAccumulationTime
func NewGenesisAccumulationTime(ctype string, prevTime time.Time) GenesisAccumulationTime {
	return GenesisAccumulationTime{
		CollateralType:           ctype,
		PreviousAccumulationTime: prevTime,
	}
}

// GenesisAccumulationTimes slice of GenesisAccumulationTime
type GenesisAccumulationTimes []GenesisAccumulationTime

// Validate performs validation of GenesisAccumulationTimes
func (gats GenesisAccumulationTimes) Validate() error {
	for _, gat := range gats {
		if err := gat.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate performs validation of GenesisAccumulationTime
func (gat GenesisAccumulationTime) Validate() error {
	if len(gat.CollateralType) == 0 {
		return fmt.Errorf("genesis accumulation time's collateral type must be defined")
	}
	return nil
}