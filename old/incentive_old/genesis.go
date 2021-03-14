package incentive

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lcnem/jpyx/x/incentive/keeper"
	"github.com/lcnem/jpyx/x/incentive/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, supplyKeeper types.SupplyKeeper, cdpKeeper types.CdpKeeper, gs types.GenesisState) {

	// check if the module account exists
	moduleAcc := supplyKeeper.GetModuleAccount(ctx, types.IncentiveMacc)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.IncentiveMacc))
	}

	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	for _, rp := range gs.Params.JPYXMintingRewardPeriods {
		_, found := cdpKeeper.GetCollateral(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("jpyx minting collateral type %s not found in cdp collateral types", rp.CollateralType))
		}
		k.SetJPYXMintingRewardFactor(ctx, rp.CollateralType, sdk.ZeroDec())
	}

	k.SetParams(ctx, gs.Params)

	for _, gat := range gs.JPYXAccumulationTimes {
		k.SetPreviousJPYXMintingAccrualTime(ctx, gat.CollateralType, gat.PreviousAccumulationTime)
	}

	for i, claim := range gs.JPYXMintingClaims {
		for j, ri := range claim.RewardIndexes {
			if ri.RewardFactor != sdk.ZeroDec() {
				gs.JPYXMintingClaims[i].RewardIndexes[j].RewardFactor = sdk.ZeroDec()
			}
		}
		k.SetJPYXMintingClaim(ctx, claim)
	}
}

// ExportGenesis export genesis state for incentive module
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	params := k.GetParams(ctx)

	jpyxClaims := k.GetAllJPYXMintingClaims(ctx)

	synchronizedJpyxClaims := types.JPYXMintingClaims{}

	for _, jpyxClaim := range jpyxClaims {
		claim, err := k.SynchronizeJPYXMintingClaim(ctx, jpyxClaim)
		if err != nil {
			panic(err)
		}
		for i := range claim.RewardIndexes {
			claim.RewardIndexes[i].RewardFactor = sdk.ZeroDec()
		}
		synchronizedJpyxClaims = append(synchronizedJpyxClaims, claim)
	}

	var jpyxMintingGats GenesisAccumulationTimes
	for _, rp := range params.JPYXMintingRewardPeriods {
		pat, found := k.GetPreviousJPYXMintingAccrualTime(ctx, rp.CollateralType)
		if !found {
			panic(fmt.Sprintf("expected previous jpyx minting reward accrual time to be set in state for %s", rp.CollateralType))
		}
		gat := types.NewGenesisAccumulationTime(rp.CollateralType, pat)
		jpyxMintingGats = append(jpyxMintingGats, gat)
	}

	return types.NewGenesisState(params, jpyxMintingGats, synchronizedJpyxClaims)
}