package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	// Methods imported from account should be defined here
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI
	GetModuleAddress(module string) sdk.AccAddress
}

type BankKeeper interface {
	// Methods imported from bank should be defined here
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, formModule string, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, amt sdk.Coins) error
}
