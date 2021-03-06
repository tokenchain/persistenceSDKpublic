package auth

import (
	"errors"

	sdk "github.com/commitHub/commitBlockchain/types"
	"github.com/commitHub/commitBlockchain/wire"
	crypto "github.com/tendermint/tendermint/crypto"
)

// Account is a standard account using a sequence number for replay protection
// and a pubkey for authentication.
type Account interface {
	GetAddress() sdk.AccAddress
	SetAddress(sdk.AccAddress) error // errors if already set.

	GetPubKey() crypto.PubKey // can return nil.
	SetPubKey(crypto.PubKey) error

	GetAccountNumber() int64
	SetAccountNumber(int64) error

	GetSequence() int64
	SetSequence(int64) error

	GetCoins() sdk.Coins
	SetCoins(sdk.Coins) error

	GetAssetPegWallet() sdk.AssetPegWallet
	SetAssetPegWallet(sdk.AssetPegWallet) error

	GetFiatPegWallet() sdk.FiatPegWallet
	SetFiatPegWallet(sdk.FiatPegWallet) error
}

// AccountDecoder unmarshals account bytes
type AccountDecoder func(accountBytes []byte) (Account, error)

//-----------------------------------------------------------
// BaseAccount

var _ Account = (*BaseAccount)(nil)

// BaseAccount - base account structure.
// Extend this by embedding this in your AppAccount.
// See the examples/basecoin/types/account.go for an example.
type BaseAccount struct {
	Address        sdk.AccAddress     `json:"address"`
	Coins          sdk.Coins          `json:"coins"`
	FiatPegWallet  sdk.FiatPegWallet  `json:"fiatPegWallet"`
	AssetPegWallet sdk.AssetPegWallet `json:"assetPegWallet"`
	PubKey         crypto.PubKey      `json:"public_key"`
	AccountNumber  int64              `json:"account_number"`
	Sequence       int64              `json:"sequence"`
}

//ProtoBaseAccount : Prototype function for BaseAccount
func ProtoBaseAccount() Account {
	return &BaseAccount{}
}

// NewBaseAccountWithAddress : returns BaseAccount with an updated address
func NewBaseAccountWithAddress(addr sdk.AccAddress) BaseAccount {
	return BaseAccount{
		Address: addr,
	}
}

// GetAddress : Implements sdk.Account.
func (acc BaseAccount) GetAddress() sdk.AccAddress {
	return acc.Address
}

// SetAddress : Implements sdk.Account.
func (acc *BaseAccount) SetAddress(addr sdk.AccAddress) error {
	if len(acc.Address) != 0 {
		return errors.New("cannot override BaseAccount address")
	}
	acc.Address = addr
	return nil
}

// GetPubKey : Implements sdk.Account.
func (acc BaseAccount) GetPubKey() crypto.PubKey {
	return acc.PubKey
}

// SetPubKey : Implements sdk.Account.
func (acc *BaseAccount) SetPubKey(pubKey crypto.PubKey) error {
	acc.PubKey = pubKey
	return nil
}

// GetCoins : Implements sdk.Account.
func (acc *BaseAccount) GetCoins() sdk.Coins {
	return acc.Coins
}

// SetCoins : Implements sdk.Account.
func (acc *BaseAccount) SetCoins(coins sdk.Coins) error {
	acc.Coins = coins
	return nil
}

//GetAssetPegWallet : getter
func (acc *BaseAccount) GetAssetPegWallet() sdk.AssetPegWallet {
	return acc.AssetPegWallet
}

//SetAssetPegWallet : setter
func (acc *BaseAccount) SetAssetPegWallet(assetPegWallet sdk.AssetPegWallet) error {
	acc.AssetPegWallet = assetPegWallet
	return nil
}

//GetFiatPegWallet : getter
func (acc *BaseAccount) GetFiatPegWallet() sdk.FiatPegWallet {
	return acc.FiatPegWallet
}

//SetFiatPegWallet : setter
func (acc *BaseAccount) SetFiatPegWallet(fiatPegWallet sdk.FiatPegWallet) error {
	acc.FiatPegWallet = fiatPegWallet
	return nil
}

// GetAccountNumber : Implements Account
func (acc *BaseAccount) GetAccountNumber() int64 {
	return acc.AccountNumber
}

// SetAccountNumber : Implements Account
func (acc *BaseAccount) SetAccountNumber(accNumber int64) error {
	acc.AccountNumber = accNumber
	return nil
}

// GetSequence : Implements sdk.Account.
func (acc *BaseAccount) GetSequence() int64 {
	return acc.Sequence
}

// SetSequence : Implements sdk.Account.
func (acc *BaseAccount) SetSequence(seq int64) error {
	acc.Sequence = seq
	return nil
}

// RegisterBaseAccount : regester base account
//----------------------------------------
// Wire
// Most users shouldn't use this, but this comes handy for tests.
func RegisterBaseAccount(cdc *wire.Codec) {
	cdc.RegisterInterface((*Account)(nil), nil)
	cdc.RegisterConcrete(&BaseAccount{}, "commit-blockchain/BaseAccount", nil)
	wire.RegisterCrypto(cdc)
}
